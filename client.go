package twitchws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"nhooyr.io/websocket"
)

type transport struct {
	Method    string `json:"method"`
	SessionID string `json:"session_id"`
}

type OnEventFn func()
type OnMessageEventFn func(Metadata, Payload)
type Option func(*Client)

type websocketMessageFn func(*Client, Metadata, []byte) (Payload, OnMessageEventFn, error)

const websocketTwitch = "wss://eventsub.wss.twitch.tv/ws"

var (
	AlreadyInUse     = errors.New("client already in use")
	NotConnected     = errors.New("client is not connected")
	ConnectionFailed = errors.New("failed to setup connection")
)

var (
	notSupported       = errors.New("message type is no supported")
	websocketReadError = errors.New("read error")
	unmarshalError     = errors.New("failed to unmarshal message")
	handlingError      = errors.New("handling error")
	connectionNotAlive = errors.New("connection is lost")
)

var (
	messageHandlers = map[string]websocketMessageFn{
		"session_welcome":   welcomeMessageHandler,
		"session_keepalive": keepaliveMessageHandler,
		"notification":      notificationMessageHandler,
		"revocation":        revocationMessageHandler,
		"session_reconnect": reconnectMessageHandler,
	}
	log = logrus.New()
)

type Metadata struct {
	MessageID        string `json:"message_id"`
	MessageType      string `json:"message_type"`
	MessageTimestamp string `json:"message_timestamp"`
}

type Payload struct {
	Payload interface{}
}

type Session struct {
	ID                      string `json:"id"`
	Status                  string `json:"status"`
	KeepaliveTimeoutSeconds int    `json:"keepalive_timeout_seconds"`
	ReconnectURL            string `json:"reconnect_url"`
	ConnectedAt             string `json:"connected_at"`
}

type Subscription struct {
	ID        string      `json:"id"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	Version   string      `json:"version"`
	Cost      string      `json:"cost"`
	Condition interface{} `json:"condition"`
	Transport transport   `json:"transport"`
	CreatedAt string      `json:"created_at"`
}

type Event struct {
	EventData interface{} `json:"event"`
}

type Client struct {
	conn         *websocket.Conn
	ctx          context.Context
	ctxCancel    context.CancelFunc
	waitGroup    *errgroup.Group
	waitGroupCtx context.Context
	workerStop   chan struct{}
	// status variables
	isActive          atomic.Bool
	isConnected       atomic.Bool
	isWelcomeReceived atomic.Bool
	msgTracking       *ttlcache.Cache[string, string]
	// connection related stuff
	url                string
	keepaliveTimeout   time.Duration
	lastHeardTimestamp time.Time
	// state event callbacks
	onConnect    OnEventFn
	onDisconnect OnEventFn
	// message event callbacks
	onWelcomeMessage      OnMessageEventFn
	onKeepaliveMessage    OnMessageEventFn
	onNotificationMessage OnMessageEventFn
	onRevocationMessage   OnMessageEventFn
	onReconnectMessage    OnMessageEventFn
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
}

func NewClientDefault(opts ...func(*Client)) *Client {
	return newClient(websocketTwitch, opts...)
}

func NewClient(url string, opts ...func(*Client)) *Client {
	return newClient(url, opts...)
}

func newClient(url string, opts ...func(*Client)) *Client {
	ctx, cancel := context.WithCancel(context.Background())

	c := &Client{
		conn:             nil,
		ctx:              ctx,
		ctxCancel:        cancel,
		keepaliveTimeout: time.Minute,
		workerStop:       make(chan struct{}, 1),
		url:              url,
		msgTracking: ttlcache.New[string, string](
			ttlcache.WithTTL[string, string](time.Second*10),
			ttlcache.WithDisableTouchOnHit[string, string](),
		),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Connect() error {
	if !c.setActive() {
		return AlreadyInUse
	}

	c.waitGroup, c.waitGroupCtx = errgroup.WithContext(c.ctx)
	c.waitGroup.Go(func() error {
		return worker(c)
	})

	return nil
}

func (c *Client) Wait() error {
	return c.waitGroup.Wait()
}

func (c *Client) Close() error {
	if !c.setInactive() {
		return NotConnected
	}

	c.workerStop <- struct{}{}
	c.ctxCancel()
	err := c.Wait()
	return err
}

func (c *Client) setActive() bool {
	return c.isActive.CompareAndSwap(false, true)
}

func (c *Client) setInactive() bool {
	return c.isActive.CompareAndSwap(true, false)
}

func (c *Client) setConnected() {
	c.isConnected.Store(true)
}

func (c *Client) setDisconnected() {
	c.isConnected.Store(false)
}

func (c *Client) getIsConnected() bool {
	return c.isConnected.Load()
}

func (c *Client) getIsWelcomeReceived() bool {
	return c.isWelcomeReceived.Load()
}

func (c *Client) isConnectionAlive() bool {
	return time.Now().Before(c.lastHeardTimestamp.Add(c.keepaliveTimeout))
}

func (c *Client) cleanUp(err error) {
	c.lastHeardTimestamp = time.Time{}
	c.isWelcomeReceived.Store(false)
	c.msgTracking.DeleteAll()

	if !c.getIsConnected() {
		return
	}

	status := websocket.StatusNormalClosure
	reason := ""

	if err != nil {
		status = websocket.StatusInternalError
		reason = fmt.Sprintf("error occurred: %s", err)
	}

	_ = c.conn.Close(status, reason)
}

func WithOnWelcome(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onWelcomeMessage = fn
	}
}

func WithOnKeepalive(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onKeepaliveMessage = fn
	}
}

func WithOnNotification(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onNotificationMessage = fn
	}
}

func WithOnRevocation(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onRevocationMessage = fn
	}
}

func WithOnReconnect(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onReconnectMessage = fn
	}
}

func worker(c *Client) error {
	var err error
	defer c.setDisconnected()

	for {
		c.conn, _, err = websocket.Dial(c.ctx, c.url, nil)

		if err != nil {
			err = errors.Join(err, ConnectionFailed)
			return err
		}

		c.setConnected()
		var shouldExit bool
		shouldExit, err = loop(c)

		if shouldExit {
			return err
		}
	}
}

func loop(c *Client) (bool, error) {
	var err error
	defer c.cleanUp(err)

	for {
		select {
		case <-c.workerStop:
			return true, nil
		default:
			err = singleMessageHandler(c)

			if err != nil {
				if errors.Is(err, context.Canceled) {
					return true, err
				} else if !errors.Is(err, notSupported) {
					log.Error(err)
					return false, err
				}
			}
		}

		c.msgTracking.DeleteExpired()
		if c.getIsWelcomeReceived() && !c.isConnectionAlive() {
			log.Debug("no keepalive/event messages - reconnect")
			return false, connectionNotAlive
		}
	}
}

func singleMessageHandler(c *Client) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.keepaliveTimeout)
	defer cancel()

	var p Payload
	msgType, data, err := c.conn.Read(ctx)

	if err != nil || msgType == websocket.MessageBinary {
		return errors.Join(err, websocketReadError)
	}

	m, err := unmarshalMetadata(data)

	if err != nil {
		return errors.Join(err, unmarshalError)
	}

	item := c.msgTracking.Get(m.MessageID)

	if item != nil {
		log.Debugf("Message ID already present: %s", item.Key())
	}

	c.msgTracking.Set(m.MessageID, m.MessageTimestamp, time.Second*c.keepaliveTimeout)

	if h, ok := messageHandlers[m.MessageType]; ok {
		var onEvent OnMessageEventFn
		p, onEvent, err = h(c, m, data)

		if err != nil {
			return errors.Join(err, handlingError)
		}

		if onEvent != nil {
			onEvent(m, p)
		}
	} else {
		log.Debugf("unknown Twitch message type: %s", m.MessageType)
	}
	return nil
}

/* Message Handlers */
func welcomeMessageHandler(c *Client, metadata Metadata, data []byte) (Payload, OnMessageEventFn, error) {
	log.Debugf("Welcome received: %s", data)
	s, err := unmarshalSession(data)
	e := Payload{
		Payload: s,
	}

	if err == nil {
		// consider reported keepalive timeout to be 80% value to be used
		// to avoid disconnection due to small drift in message delivery
		c.keepaliveTimeout = time.Duration(s.KeepaliveTimeoutSeconds*100/80) * time.Second
		c.isWelcomeReceived.Store(true)
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	return e, c.onWelcomeMessage, err
}

func keepaliveMessageHandler(c *Client, metadata Metadata, data []byte) (Payload, OnMessageEventFn, error) {
	e := Payload{
		Payload: struct{}{},
	}
	err := unmarshalEnvelope(data, &e)

	if err == nil {
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	return e, c.onKeepaliveMessage, err
}

func notificationMessageHandler(c *Client, metadata Metadata, data []byte) (Payload, OnMessageEventFn, error) {
	log.Debug("Notification received")
	e := Payload{
		Payload: struct{}{},
	}
	err := unmarshalEnvelope(data, &e)

	if err == nil {
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	return Payload{}, c.onNotificationMessage, nil
}

func revocationMessageHandler(c *Client, metadata Metadata, _ []byte) (Payload, OnMessageEventFn, error) {
	log.Debug("Revocation received")
	return Payload{}, c.onRevocationMessage, nil
}

func reconnectMessageHandler(c *Client, metadata Metadata, data []byte) (Payload, OnMessageEventFn, error) {
	s, err := unmarshalSession(data)
	e := Payload{
		Payload: s,
	}
	return e, c.onReconnectMessage, err
}

func unmarshalMetadata(data []byte) (Metadata, error) {
	var m struct {
		Metadata `json:"metadata"`
	}

	err := json.Unmarshal(data, &m)

	if err != nil {
		return Metadata{}, err
	}

	return m.Metadata, nil
}

func unmarshalEnvelope(data []byte, e any) error {
	err := json.Unmarshal(data, &e)

	if err != nil {
		log.Errorf("error: %v", err)
	}

	return err
}

func unmarshalSession(data []byte) (Session, error) {
	var payload struct {
		Payload struct {
			Session `json:"session"`
		} `json:"payload"`
	}

	if err := unmarshalEnvelope(data, &payload); err != nil {
		return payload.Payload.Session, err
	}

	return payload.Payload.Session, nil
}

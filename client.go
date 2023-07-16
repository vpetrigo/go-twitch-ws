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

type websocketMessageFn func(*Client, Metadata, []byte) (Payload, OnMessageEventFn, error)
type clientState int

const websocketTwitch = "wss://eventsub.wss.twitch.tv/ws"

var (
	ErrAlreadyInUse     = errors.New("client already in use")      // WebSocket client is already in use.
	ErrNotConnected     = errors.New("client is not connected")    // WebSocket client is not connected
	ErrConnectionFailed = errors.New("failed to setup connection") // Failed to set up WebSocket connection
)

var (
	errNotSupported           = errors.New("message type is no supported")
	errWebsocketReadError     = errors.New("read error")
	errUnmarshalError         = errors.New("failed to unmarshal message")
	errHandlingError          = errors.New("handling error")
	errConnectionNotAlive     = errors.New("connection is lost")
	errNotSupportedEvent      = errors.New("unsupported event")
	errReconnectTimeoutExpire = errors.New("reconnect awaiting timeout")
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

const (
	stateInactive = iota
	stateConnecting
	stateConnected
	stateReconnecting
	stateDisconnected
)

const defaultTTLTimeoutSec = 10

type Metadata struct {
	MessageID           string `json:"message_id"`
	MessageType         string `json:"message_type"`
	MessageTimestamp    string `json:"message_timestamp"`
	SubscriptionType    string `json:"subscription_type"`
	SubscriptionVersion string `json:"subscription_version"`
}

type Payload struct {
	Payload interface{}
}

type Session struct {
	ID                      string `json:"id"`
	Status                  string `json:"status"`
	ReconnectURL            string `json:"reconnect_url"`
	ConnectedAt             string `json:"connected_at"`
	KeepaliveTimeoutSeconds int    `json:"keepalive_timeout_seconds"`
}

type Subscription struct {
	ID        string      `json:"id"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	Version   string      `json:"version"`
	Condition interface{} `json:"condition"`
	Transport transport   `json:"transport"`
	CreatedAt string      `json:"created_at"`
	Cost      int         `json:"cost"`
}

type Notification struct {
	Subscription Subscription `json:"subscription"`
	Event        interface{}  `json:"event"`
}

type Client struct {
	ctx         context.Context
	ctxCancel   context.CancelFunc
	opCtx       context.Context
	opCtxCancel context.CancelFunc
	// connection related stuff
	conn          *websocket.Conn
	reconnectConn *websocket.Conn
	// client main routine related stuff
	waitGroup    *errgroup.Group
	waitGroupCtx context.Context
	workerStop   chan struct{}

	reconnectGroup    *errgroup.Group
	reconnectGroupCtx context.Context
	// status variables
	isActive            atomic.Bool
	isConnected         atomic.Bool
	isWelcomeReceived   atomic.Bool
	isReconnectRequired atomic.Bool
	msgTracking         *ttlcache.Cache[string, string]
	state               clientState
	// connection related stuff
	url                string
	keepaliveTimeout   time.Duration
	lastHeardTimestamp time.Time
	// event callbacks
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

func NewClientDefault(opts ...Option) *Client {
	return newClient(websocketTwitch, opts...)
}

func NewClient(url string, opts ...Option) *Client {
	return newClient(url, opts...)
}

func newClient(url string, opts ...Option) *Client {
	c := &Client{
		conn:             nil,
		keepaliveTimeout: time.Minute,
		workerStop:       make(chan struct{}, 1),
		url:              url,
		msgTracking: ttlcache.New[string, string](
			ttlcache.WithTTL[string, string](time.Second*defaultTTLTimeoutSec),
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
		return ErrAlreadyInUse
	}

	c.state = stateConnecting
	c.initMainContext()
	c.initOperationContext()
	c.waitGroup, c.waitGroupCtx = errgroup.WithContext(c.operationContext())
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
		return ErrNotConnected
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

func (c *Client) initMainContext() {
	c.ctx, c.ctxCancel = context.WithCancel(context.Background())
}

func (c *Client) initOperationContext() {
	c.opCtx, c.opCtxCancel = context.WithCancel(c.mainContext())
}

func (c *Client) mainContext() context.Context {
	return c.ctx
}

func (c *Client) operationContext() context.Context {
	return c.opCtx
}

func (c *Client) cleanUp(err error) {
	c.lastHeardTimestamp = time.Time{}
	c.isWelcomeReceived.Store(false)
	c.msgTracking.DeleteAll()

	if !c.getIsConnected() {
		return
	}

	log.Debugf("Error: %s", err)

	status := websocket.StatusNormalClosure
	reason := ""

	if err != nil {
		status = websocket.StatusInternalError
		reason = fmt.Sprintf("error occurred: %s", err)
	}

	_ = c.conn.Close(status, reason)
}

func worker(c *Client) error {
	var (
		err        error
		shouldExit bool
	)

	for {
		switch c.state {
		case stateConnecting:
			err = connectingStateHandler(c)

			if err == nil {
				c.state = stateConnected
			} else {
				shouldExit = true
				c.state = stateDisconnected
			}
		case stateConnected:
			c.setConnected()

			if c.onConnect != nil {
				c.onConnect()
			}

			shouldExit, err = connectedStateHandler(c)
			if c.isReconnectRequired.CompareAndSwap(true, false) {
				c.state = stateReconnecting

				if err != nil {
					c.state = stateDisconnected
				}
			} else {
				c.state = stateDisconnected
			}
		case stateReconnecting:
			c.initOperationContext()
			_ = c.conn.Close(websocket.StatusServiceRestart, "")
			c.conn, c.reconnectConn = c.reconnectConn, nil
			c.state = stateConnected
		case stateDisconnected:
			c.cleanUp(err)
			c.setDisconnected()

			if c.onDisconnect != nil {
				c.onDisconnect()
			}

			if !shouldExit {
				c.state = stateConnecting
			} else {
				c.state = stateInactive
			}
		case stateInactive:
			return err
		default:
			log.Errorf("unsupported state: %d", c.state)
		}
	}
}

func connectingStateHandler(c *Client) error {
	var err error
	c.conn, _, err = websocket.Dial(c.operationContext(), c.url, nil)

	if err != nil {
		err = errors.Join(err, ErrConnectionFailed)
		return err
	}

	return nil
}

func connectedStateHandler(c *Client) (bool, error) {
	for {
		select {
		case <-c.workerStop:
			return true, nil
		default:
			err := singleMessageHandler(c)

			if err != nil {
				if errors.Is(err, context.Canceled) {
					if c.isReconnectRequired.Load() {
						err := c.reconnectGroup.Wait()
						log.Debugf("Error: %s. Reconnect done", err)

						return false, err
					}

					return true, err
				} else if !errors.Is(err, errNotSupported) {
					log.Error(err)
					return false, err
				}
			}
		}

		c.msgTracking.DeleteExpired()
		if c.getIsWelcomeReceived() && !c.isConnectionAlive() {
			log.Debug("no keepalive/event messages - reconnect")
			return false, errConnectionNotAlive
		}
	}
}

func singleMessageHandler(c *Client) error {
	ctx, cancel := context.WithTimeout(c.operationContext(), c.keepaliveTimeout)
	defer cancel()

	msgType, data, err := c.conn.Read(ctx)

	if err != nil {
		return errors.Join(err, errWebsocketReadError)
	}

	m, err := getMessageMetadata(msgType, data)

	if err != nil {
		return err
	}

	item := c.msgTracking.Get(m.MessageID)

	if item != nil {
		log.Debugf("Message ID already present: %s", item.Key())
	}

	c.msgTracking.Set(m.MessageID, m.MessageTimestamp, time.Second*c.keepaliveTimeout)

	if h, ok := messageHandlers[m.MessageType]; ok {
		var (
			onEvent OnMessageEventFn
			p       Payload
		)
		p, onEvent, err = h(c, m, data)

		if err != nil {
			return errors.Join(err, errHandlingError)
		}

		if onEvent != nil {
			onEvent(m, p)
		}
	} else {
		log.Debugf("unknown Twitch message type: %s", m.MessageType)
	}

	return nil
}

func getMessageMetadata(msgType websocket.MessageType, data []byte) (Metadata, error) {
	if msgType == websocket.MessageBinary {
		return Metadata{}, errWebsocketReadError
	}

	m, err := unmarshalMetadata(data)

	if err != nil {
		return Metadata{}, errors.Join(err, errUnmarshalError)
	}

	return m, nil
}

func reconnectNewConnection(c *Client, url string) error {
	var err error
	c.url = url
	c.reconnectConn, _, err = websocket.Dial(c.mainContext(), url, nil)

	return err
}

func reconnectWaitWelcome(c *Client) error {
	var (
		err             error
		welcomeReceived bool
	)
	end := time.Now().Add(time.Minute)

	for {
		if end.After(time.Now()) {
			return errReconnectTimeoutExpire
		}

		err = func() error {
			rCtx, rCancel := context.WithTimeout(c.mainContext(), time.Minute)
			defer rCancel()

			msgType, data, err := c.reconnectConn.Read(rCtx)

			if err != nil {
				return err
			}

			m, err := getMessageMetadata(msgType, data)

			if err != nil {
				return err
			}

			if m.MessageType == "session_welcome" {
				_, _, err = welcomeMessageHandler(c, m, data)

				if err != nil {
					return err
				}

				welcomeReceived = true
			}

			return nil
		}()

		if err != nil {
			return err
		}

		if welcomeReceived {
			// cancel operation context to allow connections swap
			c.opCtxCancel()
			break
		}
	}

	return nil
}

func reconnectHandler(c *Client, url string) error {
	err := reconnectNewConnection(c, url)

	if err != nil {
		return err
	}

	return reconnectWaitWelcome(c)
}

func welcomeMessageHandler(c *Client, metadata Metadata, data []byte) (Payload, OnMessageEventFn, error) {
	s, err := unmarshalSession(data)
	e := Payload{
		Payload: s,
	}

	if err == nil {
		c.keepaliveTimeout = keepaliveIntervalCalc(s.KeepaliveTimeoutSeconds)
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
	notification, err := unmarshalNotification(data)

	if err == nil {
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	payload := Payload{
		Payload: notification,
	}

	log.Debugf("Notification: %+v", payload)

	return payload, c.onNotificationMessage, err
}

func revocationMessageHandler(c *Client, _ Metadata, _ []byte) (Payload, OnMessageEventFn, error) {
	log.Debug("Revocation received")
	return Payload{}, c.onRevocationMessage, nil
}

func reconnectMessageHandler(c *Client, _ Metadata, data []byte) (Payload, OnMessageEventFn, error) {
	s, err := unmarshalSession(data)
	e := Payload{
		Payload: s,
	}

	if err == nil {
		c.isReconnectRequired.Store(true)
		c.reconnectGroup, c.reconnectGroupCtx = errgroup.WithContext(c.ctx)
		c.reconnectGroup.Go(func() error {
			return reconnectHandler(c, s.ReconnectURL)
		})
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
		return Session{}, err
	}

	return payload.Payload.Session, nil
}

func unmarshalNotification(data []byte) (Notification, error) {
	var msg json.RawMessage
	payload := Payload{
		Payload: &msg,
	}

	if err := unmarshalEnvelope(data, &payload); err != nil {
		return Notification{}, err
	}

	var condMsg json.RawMessage
	notification := Notification{
		Subscription: Subscription{Condition: &condMsg},
		Event:        &msg,
	}

	if err := unmarshalEnvelope(msg, &notification); err != nil {
		return Notification{}, err
	}

	e, ok := eventSubTypes[notification.Subscription.Type]

	if !ok {
		err := fmt.Errorf("unsupported event: %s", notification.Subscription.Type)
		return Notification{}, errors.Join(errNotSupportedEvent, err)
	}

	event := e.MsgType
	if err := unmarshalEnvelope(msg, event); err != nil {
		return Notification{}, err
	}

	condition := e.ConditionType
	if err := unmarshalEnvelope(condMsg, condition); err != nil {
		return Notification{}, err
	}

	notification.Event = event
	notification.Subscription.Condition = condition
	return notification, nil
}

func keepaliveIntervalCalc(keepaliveInterval int) time.Duration {
	const keepalivePercent = 80
	// consider reported keepalive timeout to be 80% value to be used
	// to avoid disconnection due to small drift in message delivery
	return time.Duration(keepaliveInterval*100/keepalivePercent) * time.Second
}

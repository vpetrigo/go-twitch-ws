// Package twitchws provides a client library for interacting with Twitch's EventSub WebSocket API.
// This package manages WebSocket connections, handles reconnections, processes different message types,
// and exposes an interface to listen for events and notifications from Twitch EventSub.
//
// # Features
//
// • Establish and manage a WebSocket connection to the Twitch EventSub service.
//
// • Support for reconnection with automatic handling of session re-establishment.
//
// • Handlers for various Twitch EventSub message types, including session updates, notifications, keepalives, and more.
//
// • Customizable message and event handling through callback functions.
//
// • Context-driven lifecycle management, ensuring clean connection termination and resource cleanup.
//
// • Built-in support for message caching and TTL-based handling of message metadata.
//
// # Usage Example
//
// Below is an example of using the `twitchws` package to connect to the Twitch EventSub WebSocket
// and listen to event notifications.
//
//	package main
//
//	import (
//		"context"
//		"fmt"
//
//		"github.com/vpetrigo/go-twitch-ws"
//	)
//
//	const websocketTwitchTestServer = "ws://127.0.0.1:8080/ws"
//
//	func main() {
//		messageHandler := func(m *twitchws.Metadata, p *twitchws.Payload) {
//			fmt.Printf("Metadata: %+v\n", m)
//			fmt.Printf("Payload: %+v\n", p)
//		}
//		stateHandler := func(state string) func() {
//			return func() {
//				fmt.Printf("Event: %s\n", state)
//			}
//		}
//		c := twitchws.NewClient(
//			websocketTwitchTestServer,
//			twitchws.WithOnWelcome(messageHandler),
//			twitchws.WithOnNotification(messageHandler),
//			twitchws.WithOnConnect(stateHandler("Connect")),
//			twitchws.WithOnDisconnect(stateHandler("Disconnect")),
//			twitchws.WithOnRevocation(messageHandler))
//		// Start the WebSocket connection.
//		err := c.Connect()
//
//		if err != nil {
//			fmt.Println(err)
//		}
//		// Wait or run other application logic here.
//		err = c.Wait()
//
//		if err != nil {
//			fmt.Println(err)
//		}
//		// Properly stop the client when shutting down.
//		err = c.Close()
//
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
package twitchws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
	"github.com/jellydator/ttlcache/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// OnEventFn defines a callback function to be executed on specific client events such as connection or disconnection.
type OnEventFn func()

// OnMessageEventFn defines a callback function to process message events, receiving metadata and payload as parameters.
type OnMessageEventFn func(*Metadata, *Payload)

// websocketMessageFn defines a function type that processes a websocket message, returning a Payload, event callback, and error.
type websocketMessageFn func(*Client, *Metadata, []byte) (*Payload, OnMessageEventFn, error)

// clientState represents the internal state of a client, capturing various states during its lifecycle and operations.
type clientState int

// websocketTwitch is the WebSocket endpoint URL for connecting to Twitch EventSub services.
const websocketTwitch = "wss://eventsub.wss.twitch.tv/ws"

var (
	ErrAlreadyInUse     = errors.New("client already in use")      // WebSocket client is already in use.
	ErrNotConnected     = errors.New("client is not connected")    // WebSocket client is not connected
	ErrConnectionFailed = errors.New("failed to setup connection") // Failed to set up WebSocket connection
)

var (

	// errNotSupported indicates that the message type provided is not supported by the application.
	errNotSupported = errors.New("message type is no supported")

	// errWebsocketReadError represents an error that occurs during reading from a WebSocket connection.
	errWebsocketReadError = errors.New("read error")

	// errUnmarshalError represents an error indicating failure to unmarshal a message from its data structure.
	errUnmarshalError = errors.New("failed to unmarshal message")

	// errHandlingError represents an error that occurs during the handling of a message or operation.
	errHandlingError = errors.New("handling error")

	// errConnectionNotAlive indicates that the connection to the server is no longer active or has been lost.
	errConnectionNotAlive = errors.New("connection is lost")

	// errNotSupportedEvent represents an error indicating that an attempted operation involves an unsupported event type.
	errNotSupportedEvent = errors.New("unsupported event")

	// errReconnectTimeoutExpire represents an error indicating that the reconnection process exceeded the allowed timeout.
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

	// stateInactive represents the client state where it is not active or engaged in any connection-related process.
	stateInactive = iota

	// stateConnecting represents the client state during the process of establishing a connection.
	stateConnecting

	// stateConnected indicates that the client has successfully established a connection and is in a stable connected state.
	stateConnected

	// stateReconnecting indicates that the client is attempting to reconnect after being disconnected.
	stateReconnecting

	// stateDisconnected represents the state where the client has been disconnected and requires cleanup or reconnection setup.
	stateDisconnected
)

// defaultTTLTimeoutSec defines the default time-to-live duration in seconds for cached messages in the client's tracking system.
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

type EventsubSubscription struct {
	ID        string            `json:"id"`
	Status    string            `json:"status"`
	Type      string            `json:"type"`
	Version   string            `json:"version"`
	Condition EventsubCondition `json:"condition"`
	Transport EventsubTransport `json:"transport"`
	CreatedAt string            `json:"created_at"`
	Cost      int64             `json:"cost"`
}

type EventsubTransport struct {
	Method    string `json:"method"`
	Callback  string `json:"callback,omitempty"`
	SessionID string `json:"session_id,omitempty"`
}

type EventsubCondition struct {
	BroadcasterUserID     string `json:"broadcaster_user_id,omitempty"`
	ToBroadcasterUserID   string `json:"to_broadcaster_user_id,omitempty"`
	UserID                string `json:"user_id,omitempty"`
	FromBroadcasterUserID string `json:"from_broadcaster_user_id,omitempty"`
	ModeratorUserID       string `json:"moderator_user_id,omitempty"`
	ClientID              string `json:"client_id,omitempty"`
	ExtensionClientID     string `json:"extension_client_id,omitempty"`
	OrganizationID        string `json:"organization_id,omitempty"`
	CategoryID            string `json:"category_id,omitempty"`
	CampaignID            string `json:"campaign_id,omitempty"`
}

type Notification struct {
	Subscription EventsubSubscription `json:"subscription"`
	Event        interface{}          `json:"event"`
}

type Client struct {
	// ctx is the main context for the client, used to manage the lifecycle and cancelation of ongoing operations.
	ctx context.Context

	// ctxCancel is a context cancel function used to terminate the main context of the client.
	ctxCancel context.CancelFunc

	// opCtx is the operational context used for managing and canceling ongoing operations within the client lifecycle.
	opCtx context.Context

	// opCtxCancel defines a cancel function for the operational context, used to terminate operations gracefully.
	opCtxCancel context.CancelFunc

	// conn represents the active WebSocket connection used for communication between the client and the server.
	conn *websocket.Conn

	// reconnectConn stores a websocket connection used for reconnecting after the main connection is lost or closed.
	reconnectConn *websocket.Conn

	// waitGroup is used to manage a group of goroutines and wait for their completion or capture their errors collectively.
	waitGroup *errgroup.Group

	// waitGroupCtx is the context associated with the wait group, enabling cancellation and shared context across routines.
	waitGroupCtx context.Context

	// workerStop is a channel used to signal the worker goroutine to stop its execution.
	workerStop chan struct{}

	// reconnectGroup manages a group of goroutines for handling reconnection logic within the client lifecycle.
	reconnectGroup *errgroup.Group

	// reconnectGroupCtx provides the context for managing the lifecycle of the reconnection goroutine(s) within the client.
	reconnectGroupCtx context.Context

	// isActive is an atomic.Boolean that indicates whether the client is currently active or in use.
	isActive atomic.Bool

	// isConnected indicates whether the client is currently connected. Managed using an atomic boolean for concurrency safety.
	isConnected atomic.Bool

	// isWelcomeReceived indicates whether the welcome message from the server has been successfully received and processed.
	isWelcomeReceived atomic.Bool

	// isReconnectRequired indicates whether the client is required to reconnect, controlled via atomic operations.
	isReconnectRequired atomic.Bool

	// msgTracking maintains a cache for tracking message IDs along with their timestamps to handle deduplication and expiration.
	msgTracking *ttlcache.Cache[string, string]

	// state represents the current lifecycle state of the Client, determining its operational mode and transitions.
	state clientState

	// url specifies the WebSocket server address the client connects to or interacts with.
	url string

	// keepaliveTimeout represents the duration within which a keepalive message is expected to maintain connection health.
	keepaliveTimeout time.Duration

	// lastHeardTimestamp stores the last known timestamp when the client received a message or activity.
	lastHeardTimestamp time.Time

	// onConnect is a callback executed when the client successfully connects to the WebSocket server.
	onConnect OnEventFn

	// onDisconnect is the callback function executed when the client disconnects from the server.
	onDisconnect OnEventFn

	// onWelcomeMessage is a callback function to handle events triggered on receiving a welcome message from the server.
	onWelcomeMessage OnMessageEventFn

	// onKeepaliveMessage is a callback function invoked to handle keepalive messages received by the client.
	onKeepaliveMessage OnMessageEventFn

	// onNotificationMessage defines a callback for handling incoming notification messages with associated metadata and payload.
	onNotificationMessage OnMessageEventFn

	// onRevocationMessage is a callback function triggered when a revocation message is received by the client.
	onRevocationMessage OnMessageEventFn

	// onReconnectMessage defines a callback function triggered when a reconnect message is received by the client.
	onReconnectMessage OnMessageEventFn
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
}

// NewClientDefault creates a new Client instance with the default websocketTwitch URL and optional configuration options.
func NewClientDefault(opts ...Option) *Client {
	return newClient(websocketTwitch, opts...)
}

// NewClient creates and returns a new Client instance configured with the provided WebSocket URL and optional settings.
func NewClient(url string, opts ...Option) *Client {
	return newClient(url, opts...)
}

// newClient creates and initializes a new Client instance with the specified URL and optional configuration options.
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

// Connect establishes a connection by initializing contexts, transitioning to the connecting state, and starting the worker.
// Returns an error if the client is already active.
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

// Wait blocks until all client tasks have completed. Returns any error encountered during awaiting.
func (c *Client) Wait() error {
	return c.waitGroup.Wait()
}

// Close gracefully terminates the client's connection, stops the worker, cancels contexts, and waits for cleanup to complete.
func (c *Client) Close() error {
	if !c.setInactive() {
		return ErrNotConnected
	}

	c.workerStop <- struct{}{}
	c.ctxCancel()
	err := c.Wait()
	return err
}

// setActive attempts to set the client's active status to true using an atomic operation.
// Returns true if the status was successfully changed from false to true, false otherwise.
func (c *Client) setActive() bool {
	return c.isActive.CompareAndSwap(false, true)
}

// setInactive attempts to set the client's active status to false using an atomic operation.
// Returns true if the status was successfully changed from true to false, false otherwise.
func (c *Client) setInactive() bool {
	return c.isActive.CompareAndSwap(true, false)
}

// setConnected sets the client's connected status to true using an atomic operation.
func (c *Client) setConnected() {
	c.isConnected.Store(true)
}

// setDisconnected sets the client's connected status to false using an atomic operation.
func (c *Client) setDisconnected() {
	c.isConnected.Store(false)
}

// getIsConnected returns the current connection status of the client as a boolean value.
func (c *Client) getIsConnected() bool {
	return c.isConnected.Load()
}

// getIsWelcomeReceived returns the current state of whether a welcome message has been received by the client.
func (c *Client) getIsWelcomeReceived() bool {
	return c.isWelcomeReceived.Load()
}

// isConnectionAlive checks if the connection is still alive by comparing the current time with the last heard timestamp.
func (c *Client) isConnectionAlive() bool {
	return time.Now().Before(c.lastHeardTimestamp.Add(c.keepaliveTimeout))
}

// initMainContext initializes the main context and its cancellation function for the Client.
func (c *Client) initMainContext() {
	c.ctx, c.ctxCancel = context.WithCancel(context.Background())
}

// initOperationContext initializes the operation context and its cancellation function using the main context.
func (c *Client) initOperationContext() {
	c.opCtx, c.opCtxCancel = context.WithCancel(c.mainContext())
}

// mainContext returns the main context associated with the client instance. It is used to manage overall context lifecycles.
func (c *Client) mainContext() context.Context {
	return c.ctx
}

// operationContext returns the operation-specific context for managing tasks and cancellation within the Client instance.
func (c *Client) operationContext() context.Context {
	return c.opCtx
}

// cleanUp resets client state, clears message tracking, and closes the connection with appropriate status and reason.
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

// worker manages the state transitions of the Client, handling connection, reconnection, and disconnection processes.
// Returns an error if encountered during state handling or cleanup.
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

// connectingStateHandler attempts to establish a WebSocket connection for the provided client.
// Returns an error if the connection fails, appending ErrConnectionFailed to the error chain.
func connectingStateHandler(c *Client) error {
	var err error
	c.conn, _, err = websocket.Dial(c.operationContext(), c.url, nil)

	if err != nil {
		err = errors.Join(err, ErrConnectionFailed)
		return err
	}

	return nil
}

// connectedStateHandler manages the state of the connected client, processing messages and checking connection health.
// It handles reconnections, message expiration, and client state transitions.
// Returns true if the connection should close cleanly or false if reconnection is required, along with any error encountered.
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

// singleMessageHandler processes a single incoming WebSocket message, updates message tracking, and invokes appropriate handlers.
// Returns an error if message reading, metadata extraction, or handling fails.
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
			p       *Payload
		)
		p, onEvent, err = h(c, m, data)

		if err != nil {
			return errors.Join(err, errHandlingError)
		}

		if onEvent != nil {
			onEvent(m, p)
		}
	} else {
		log.Warnf("unknown Twitch message type: %s", m.MessageType)
	}

	return nil
}

// getMessageMetadata extracts and unmarshals the metadata from a WebSocket message, returning it or an appropriate error.
func getMessageMetadata(msgType websocket.MessageType, data []byte) (*Metadata, error) {
	if msgType == websocket.MessageBinary {
		return nil, errWebsocketReadError
	}

	m, err := unmarshalMetadata(data)

	if err != nil {
		return nil, errors.Join(err, errUnmarshalError)
	}

	return m, nil
}

// reconnectNewConnection attempts to reconnect the client to a new WebSocket connection using the specified URL.
// It updates the client's URL and establishes a new connection with the server, returning an error if it fails.
func reconnectNewConnection(c *Client, url string) error {
	var err error
	c.url = url
	c.reconnectConn, _, err = websocket.Dial(c.mainContext(), url, nil)

	return err
}

// reconnectWaitWelcome attempts to read and process a "session_welcome" message from the reconnect connection.
// It waits for the welcome message within a timeout duration and transitions the client if successful.
// Returns an error if the timeout expires or if any issues occur during the process.
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

// reconnectHandler attempts to reconnect the client to a new connection and waits for a "session_welcome" message.
// Returns an error if reconnecting or receiving the welcome message fails.
func reconnectHandler(c *Client, url string) error {
	err := reconnectNewConnection(c, url)

	if err != nil {
		return err
	}

	return reconnectWaitWelcome(c)
}

// welcomeMessageHandler processes the "session_welcome" message, updates client state, and returns payload and callback.
func welcomeMessageHandler(c *Client, metadata *Metadata, data []byte) (*Payload, OnMessageEventFn, error) {
	s, err := unmarshalSession(data)
	e := Payload{
		Payload: s,
	}

	if err == nil {
		c.keepaliveTimeout = keepaliveIntervalCalc(s.KeepaliveTimeoutSeconds)
		c.isWelcomeReceived.Store(true)
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	return &e, c.onWelcomeMessage, err
}

// keepaliveMessageHandler processes "session_keepalive" messages, updates last heard timestamp, and returns payload and callback.
func keepaliveMessageHandler(c *Client, metadata *Metadata, data []byte) (*Payload, OnMessageEventFn, error) {
	e := Payload{
		Payload: struct{}{},
	}
	err := unmarshalEnvelope(data, &e)

	if err == nil {
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	return &e, c.onKeepaliveMessage, err
}

// notificationMessageHandler processes "notification" messages by parsing data into a payload and updating client state.
// It returns the parsed payload, the onNotificationMessage callback function, and any error encountered during processing.
func notificationMessageHandler(c *Client, metadata *Metadata, data []byte) (*Payload, OnMessageEventFn, error) {
	payload, err := processNotification(data)

	if err == nil {
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	log.Debugf("Notification: %+v", payload)

	return payload, c.onNotificationMessage, err
}

// revocationMessageHandler processes a "revocation" message, updates the client's state, and returns payload and callback.
func revocationMessageHandler(c *Client, metadata *Metadata, data []byte) (*Payload, OnMessageEventFn, error) {
	payload, err := processNotification(data)

	if err == nil {
		c.lastHeardTimestamp, err = time.Parse(time.RFC3339Nano, metadata.MessageTimestamp)
	}

	log.Debugf("Revocation received: %#v", payload)

	return payload, c.onRevocationMessage, err
}

// reconnectMessageHandler processes the "session_reconnect" message, initializing reconnect tasks and callbacks.
func reconnectMessageHandler(c *Client, _ *Metadata, data []byte) (*Payload, OnMessageEventFn, error) {
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

	return &e, c.onReconnectMessage, err
}

// unmarshalMetadata unmarshals JSON-encoded data into a Metadata struct and returns it or an error if unmarshaling fails.
func unmarshalMetadata(data []byte) (*Metadata, error) {
	var m struct {
		Metadata `json:"metadata"`
	}

	err := json.Unmarshal(data, &m)

	if err != nil {
		return nil, err
	}

	return &m.Metadata, nil
}

// unmarshalEnvelope deserializes JSON data into the provided interface and logs any errors encountered during unmarshalling.
func unmarshalEnvelope(data []byte, e any) error {
	err := json.Unmarshal(data, &e)

	if err != nil {
		log.Errorf("error: %v", err)
	}

	return err
}

// unmarshalSession extracts a Session object from a JSON byte slice, returning it or an error if deserialization fails.
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

// unmarshalNotification parses the provided JSON data into a Notification object and validates its subscription and event details.
func unmarshalNotification(data []byte) (Notification, error) {
	var msg json.RawMessage
	payload := Payload{
		Payload: &msg,
	}

	if err := unmarshalEnvelope(data, &payload); err != nil {
		return Notification{}, err
	}

	notification := Notification{
		Subscription: EventsubSubscription{},
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

	var foundEventScope *eventSubScope

	for i, v := range e {
		if v.Version == notification.Subscription.Version {
			foundEventScope = &e[i]
			break
		}
	}

	if foundEventScope == nil {
		err := fmt.Errorf("unsupported event version: %s", notification.Subscription.Version)
		return Notification{}, errors.Join(errNotSupportedEvent, err)
	}

	if foundEventScope.MsgType == nil {
		err := fmt.Errorf("unsupported message: %s", notification.Subscription.Type)
		return Notification{}, errors.Join(errNotSupportedEvent, err)
	}

	event := foundEventScope.MsgType
	if err := unmarshalEnvelope(msg, event); err != nil {
		return Notification{}, err
	}

	notification.Event = event

	return notification, nil
}

// keepaliveIntervalCalc calculates the keepalive timeout duration based on a given interval reduced by a predefined percentage.
func keepaliveIntervalCalc(keepaliveInterval int) time.Duration {
	const keepalivePercent = 80
	// consider reported keepalive timeout to be 80% value to be used
	// to avoid disconnection due to small drift in message delivery
	return time.Duration(keepaliveInterval*100/keepalivePercent) * time.Second
}

// processNotification parses the notification data into a Payload and returns it along with any error encountered.
func processNotification(data []byte) (*Payload, error) {
	notification, err := unmarshalNotification(data)

	payload := Payload{
		Payload: notification,
	}

	return &payload, err
}

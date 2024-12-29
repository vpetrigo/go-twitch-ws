package twitchws

// Option is a functional option used to configure a Client instance.
type Option func(*Client)

// WithOnWelcome sets a callback function to handle welcome message events for the client.
func WithOnWelcome(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onWelcomeMessage = fn
	}
}

// WithOnKeepalive sets a callback function to handle keepalive messages for the client instance.
func WithOnKeepalive(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onKeepaliveMessage = fn
	}
}

// WithOnNotification sets the callback function to handle notification messages received by the client.
func WithOnNotification(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onNotificationMessage = fn
	}
}

// WithOnRevocation sets a callback function to handle revocation messages and returns an Option for the Client configuration.
func WithOnRevocation(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onRevocationMessage = fn
	}
}

// WithOnReconnect configures a callback function to be invoked when a reconnect message is received by the client.
func WithOnReconnect(fn OnMessageEventFn) Option {
	return func(c *Client) {
		c.onReconnectMessage = fn
	}
}

// WithOnConnect sets the callback function to be executed when the client successfully connects to the WebSocket server.
func WithOnConnect(fn OnEventFn) Option {
	return func(c *Client) {
		c.onConnect = fn
	}
}

// WithOnDisconnect sets the callback function to be executed when the client disconnects.
func WithOnDisconnect(fn OnEventFn) Option {
	return func(c *Client) {
		c.onDisconnect = fn
	}
}

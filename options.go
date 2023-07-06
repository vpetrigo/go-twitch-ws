package twitchws

type Option func(*Client)

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

func WithOnConnect(fn OnEventFn) Option {
	return func(c *Client) {
		c.onConnect = fn
	}
}

func WithOnDisconnect(fn OnEventFn) Option {
	return func(c *Client) {
		c.onDisconnect = fn
	}
}

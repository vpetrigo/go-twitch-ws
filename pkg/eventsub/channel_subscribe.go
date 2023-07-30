package eventsub

type ChannelSubscribeEvent struct {
	UserID               string `json:"user_id"`                // The user ID for the user who subscribed to the specified channel.
	UserLogin            string `json:"user_login"`             // The user login for the user who subscribed to the specified channel.
	UserName             string `json:"user_name"`              // The user display name for the user who subscribed to the specified channel.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Tier                 string `json:"tier"`                   // The tier of the subscription.
	IsGift               bool   `json:"is_gift"`                // Whether the subscription is a gift.
}

type ChannelSubscribeEventCondition struct{}

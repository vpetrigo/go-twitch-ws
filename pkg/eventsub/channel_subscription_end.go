package eventsub

type ChannelSubscriptionEndEvent struct {
	UserID               string `json:"user_id"`                // The user ID for the user whose subscription ended.
	UserLogin            string `json:"user_login"`             // The user login for the user whose subscription ended.
	UserName             string `json:"user_name"`              // The user display name for the user whose subscription ended.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster display name.
	Tier                 string `json:"tier"`                   // The tier of the subscription that ended.
	IsGift               bool   `json:"is_gift"`                // Whether the subscription was a gift.
}

type ChannelSubscriptionEndEventCondition struct{}

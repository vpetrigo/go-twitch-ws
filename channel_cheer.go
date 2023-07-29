package twitchws

type ChannelCheerEvent struct {
	IsAnonymous          bool   `json:"is_anonymous"`           // Whether the user cheered anonymously or not.
	UserId               string `json:"user_id"`                // The user ID for the user who cheered on the specified channel.
	UserLogin            string `json:"user_login"`             // The user login for the user who cheered on the specified channel.
	UserName             string `json:"user_name"`              // The user display name for the user who cheered on the specified channel.
	BroadcasterUserId    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Message              string `json:"message"`                // The message sent with the cheer.
	Bits                 int    `json:"bits"`                   // The number of bits cheered.
}

type ChannelCheerEventCondition struct{}

package eventsub

type ChannelFollowEvent struct {
	UserID               string `json:"user_id"`                // The user ID for the user now following the specified channel.
	UserLogin            string `json:"user_login"`             // The user login for the user now following the specified channel.
	UserName             string `json:"user_name"`              // The user display name for the user now following the specified channel.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	FollowedAt           string `json:"followed_at"`            // RFC3339 timestamp of when the follow occurred.
}

type ChannelFollowEventCondition struct{}

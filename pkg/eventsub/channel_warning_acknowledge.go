package eventsub

type ChannelWarningAcknowledgeEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The user ID of the broadcaster.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The login of the broadcaster.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The user name of the broadcaster.
	UserID               string `json:"user_id"`                // The ID of the user that has acknowledged their warning.
	UserLogin            string `json:"user_login"`             // The login of the user that has acknowledged their warning.
	UserName             string `json:"user_name"`              // The user name of the user that has acknowledged their warning.
}

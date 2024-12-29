package eventsub

type ChannelChatClearUserMessagesEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster login.
	TargetUserID         string `json:"target_user_id"`         // The ID of the user that was banned or put in a timeout.
	TargetUserName       string `json:"target_user_name"`       // The user name of the user that was banned or put in a timeout.
	TargetUserLogin      string `json:"target_user_login"`      // The user login of the user that was banned or put in a timeout.
}

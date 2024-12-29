package eventsub

type ChannelChatClearEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster login.
}

package eventsub

type ChannelChatMessageDeleteEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster login.
	TargetUserID         string `json:"target_user_id"`         // The ID of the user whose message was deleted.
	TargetUserName       string `json:"target_user_name"`       // The user name of the user whose message was deleted.
	TargetUserLogin      string `json:"target_user_login"`      // The user login of the user whose message was deleted.
	MessageID            string `json:"message_id"`             // A UUID that identifies the message that was removed.
}

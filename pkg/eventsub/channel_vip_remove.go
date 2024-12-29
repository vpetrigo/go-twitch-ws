package eventsub

type ChannelVIPRemoveEvent struct {
	UserID               string `json:"user_id"`                // The ID of the user who was removed as a VIP.
	UserLogin            string `json:"user_login"`             // The login of the user who was removed as a VIP.
	UserName             string `json:"user_name"`              // The display name of the user who was removed as a VIP.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The ID of the broadcaster.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The login of the broadcaster.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The display name of the broadcaster.
}

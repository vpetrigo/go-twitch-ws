package eventsub

type StreamOfflineEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster’s user id.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s user login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s user display name.
}

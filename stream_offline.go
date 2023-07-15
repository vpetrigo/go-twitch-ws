package twitchws

// StreamOfflineEvent Contains the broadcaster user ID and broadcaster user name.
//
// **Authorization:** not required.
type StreamOfflineEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster’s user ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s user login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s user display name.
}

type StreamOfflineCondition struct {
	BroadcasterUserID string `json:"broadcaster_user_id"` // The broadcaster user ID you want to get stream offline notifications for.
}

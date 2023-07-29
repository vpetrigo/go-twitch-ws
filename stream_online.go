package twitchws

// StreamOnlineEvent Contains the stream ID, broadcaster user ID,
// broadcaster username, and the stream type.
//
// **Authorization:** not required.
type StreamOnlineEvent struct {
	ID                   string `json:"id"`                     // id of the steam.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster’s user id.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s user login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s user display name.
	Type                 string `json:"type"`                   // The stream type. Valid values are: `live`, `playlist`, `watch_party`, `premiere`, `rerun`.
	StartedAt            string `json:"started_at"`             // The timestamp at which the stream went online at.
}

type StreamOnlineEventCondition struct {
	BroadcasterUserID string `json:"broadcaster_user_id"` // The broadcaster user ID you want to get stream online notifications for.
}

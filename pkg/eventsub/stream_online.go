package eventsub

type StreamOnlineEvent struct {
	ID                   string `json:"id"`                     // The id of the stream.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster’s user id.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s user login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s user display name.
	Type                 string `json:"type"`                   // The stream type.
	StartedAt            string `json:"started_at"`             // The timestamp at which the stream went online at.
}

package eventsub

type ShoutoutReceivedEvent struct {
	BroadcasterUserID        string `json:"broadcaster_user_id"`         // An ID that identifies the broadcaster that received the Shoutout.
	BroadcasterUserLogin     string `json:"broadcaster_user_login"`      // The broadcaster’s login name.
	BroadcasterUserName      string `json:"broadcaster_user_name"`       // The broadcaster’s display name.
	FromBroadcasterUserID    string `json:"from_broadcaster_user_id"`    // An ID that identifies the broadcaster that sent the Shoutout.
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"` // The broadcaster’s login name.
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`  // The broadcaster’s display name.
	ViewerCount              int    `json:"viewer_count"`                // The number of users that were watching the from-broadcaster’s stream at the time of the Shoutout.
	StartedAt                string `json:"started_at"`                  // The UTC timestamp (in RFC3339 format) of when the moderator sent the Shoutout.
}

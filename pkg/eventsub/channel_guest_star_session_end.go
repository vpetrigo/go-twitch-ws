package eventsub

type ChannelGuestStarSessionEndEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The non-host broadcaster user ID.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The non-host broadcaster display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The non-host broadcaster login.
	SessionID            string `json:"session_id"`             // ID representing the unique session that was started.
	StartedAt            string `json:"started_at"`             // RFC3339 timestamp indicating the time the session began.
	EndedAt              string `json:"ended_at"`               // RFC3339 timestamp indicating the time the session ended.
	HostUserID           string `json:"host_user_id"`           // User ID of the host channel.
	HostUserName         string `json:"host_user_name"`         // The host display name.
	HostUserLogin        string `json:"host_user_login"`        // The host login.
}

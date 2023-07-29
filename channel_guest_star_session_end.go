package twitchws

type ChannelGuestStarSessionEndEvent struct {
	BroadcasterUserId    string `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster login.
	SessionId            string `json:"session_id"`             // ID representing the unique session that was started.
	StartedAt            string `json:"started_at"`             // RFC3339 timestamp indicating the time the session began.
	EndedAt              string `json:"ended_at"`               // RFC3339 timestamp indicating the time the session ended.
}

type ChannelGuestStarSessionEndEventCondition struct{}

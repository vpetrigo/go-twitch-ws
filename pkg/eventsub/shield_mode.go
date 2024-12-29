package eventsub

type ShieldModeEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // An ID that identifies the broadcaster whose Shield Mode status was updated.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s login name.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s display name.
	ModeratorUserID      string `json:"moderator_user_id"`      // An ID that identifies the moderator that updated the Shield Mode’s status.
	ModeratorUserLogin   string `json:"moderator_user_login"`   // The moderator’s login name.
	ModeratorUserName    string `json:"moderator_user_name"`    // The moderator’s display name.
	StartedAt            string `json:"started_at"`             // The UTC timestamp (in RFC3339 format) of when the moderator activated Shield Mode.
	EndedAt              string `json:"ended_at"`               // The UTC timestamp (in RFC3339 format) of when the moderator deactivated Shield Mode.
}

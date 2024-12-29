package eventsub

type ShoutoutCreateEvent struct {
	BroadcasterUserID      string `json:"broadcaster_user_id"`       // An ID that identifies the broadcaster that sent the Shoutout.
	BroadcasterUserLogin   string `json:"broadcaster_user_login"`    // The broadcaster’s login name.
	BroadcasterUserName    string `json:"broadcaster_user_name"`     // The broadcaster’s display name.
	ToBroadcasterUserID    string `json:"to_broadcaster_user_id"`    // An ID that identifies the broadcaster that received the Shoutout.
	ToBroadcasterUserLogin string `json:"to_broadcaster_user_login"` // The broadcaster’s login name.
	ToBroadcasterUserName  string `json:"to_broadcaster_user_name"`  // The broadcaster’s display name.
	ModeratorUserID        string `json:"moderator_user_id"`         // An ID that identifies the moderator that sent the Shoutout.
	ModeratorUserLogin     string `json:"moderator_user_login"`      // The moderator’s login name.
	ModeratorUserName      string `json:"moderator_user_name"`       // The moderator’s display name.
	ViewerCount            int    `json:"viewer_count"`              // The number of users that were watching the broadcaster’s stream at the time of the Shoutout.
	StartedAt              string `json:"started_at"`                // The UTC timestamp (in RFC3339 format) of when the moderator sent the Shoutout.
	CooldownEndsAt         string `json:"cooldown_ends_at"`          // The UTC timestamp (in RFC3339 format) of when the broadcaster may send a Shoutout to a different broadcaster.
	TargetCooldownEndsAt   string `json:"target_cooldown_ends_at"`   // The UTC timestamp (in RFC3339 format) of when the broadcaster may send another Shoutout to the broadcaster in to_broadcaster_user_id.
}

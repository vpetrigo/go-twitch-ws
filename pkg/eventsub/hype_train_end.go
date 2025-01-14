package eventsub

type HypeTrainEndEvent struct {
	ID                   string             `json:"id"`                     // The Hype Train ID.
	BroadcasterUserID    string             `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string             `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string             `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Level                int                `json:"level"`                  // The final level of the Hype Train.
	Total                int                `json:"total"`                  // Total points contributed to the Hype Train.
	TopContributions     []TopContributions `json:"top_contributions"`      // The contributors with the most points contributed.
	StartedAt            string             `json:"started_at"`             // The time when the Hype Train started.
	EndedAt              string             `json:"ended_at"`               // The time when the Hype Train ended.
	CooldownEndsAt       string             `json:"cooldown_ends_at"`       // The time when the Hype Train cooldown ends so that the next Hype Train can start.
	IsGoldenKappaTrain   bool               `json:"is_golden_kappa_train"`  // Indicates if the Hype Train is a Golden Kappa Train.
}

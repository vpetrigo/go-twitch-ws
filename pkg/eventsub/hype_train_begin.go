package eventsub

type HypeTrainBeginEvent struct {
	ID                   string             `json:"id"`                     // The Hype Train ID.
	BroadcasterUserID    string             `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string             `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string             `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Total                int                `json:"total"`                  // Total points contributed to the Hype Train.
	Progress             int                `json:"progress"`               // The number of points contributed to the Hype Train at the current level.
	Goal                 int                `json:"goal"`                   // The number of points required to reach the next level.
	TopContributions     []TopContributions `json:"top_contributions"`      // The contributors with the most points contributed.
	LastContribution     LastContribution   `json:"last_contribution"`      // The most recent contribution.
	Level                int                `json:"level"`                  // The starting level of the Hype Train.
	StartedAt            string             `json:"started_at"`             // The time when the Hype Train started.
	ExpiresAt            string             `json:"expires_at"`             // The time when the Hype Train expires.
	IsGoldenKappaTrain   bool               `json:"is_golden_kappa_train"`  // Indicates if the Hype Train is a Golden Kappa Train.
}

package eventsub

type ChannelPredictionLockEvent struct {
	ID                   string     `json:"id"`                     // Channel Points Prediction ID.
	BroadcasterUserID    string     `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string     `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string     `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Title                string     `json:"title"`                  // Title for the Channel Points Prediction.
	Outcomes             []Outcomes `json:"outcomes"`               // An array of outcomes for the Channel Points Prediction.
	StartedAt            string     `json:"started_at"`             // The time the Channel Points Prediction started.
	LockedAt             string     `json:"locked_at"`              // The time the Channel Points Prediction was locked.
}

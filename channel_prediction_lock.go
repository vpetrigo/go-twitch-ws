package twitchws

type ChannelPredictionLockEvent struct {
	Id                   string      `json:"id"`                     // Channel Points Prediction ID.
	BroadcasterUserId    string      `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string      `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string      `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Title                string      `json:"title"`                  // Title for the Channel Points Prediction.
	Outcomes             interface{} `json:"outcomes"`               // An array of outcomes for the Channel Points Prediction.
	StartedAt            string      `json:"started_at"`             // The time the Channel Points Prediction started.
	LockedAt             string      `json:"locked_at"`              // The time the Channel Points Prediction was locked.
}

type ChannelPredictionLockEventCondition struct{}
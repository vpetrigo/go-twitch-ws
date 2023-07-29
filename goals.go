package twitchws

type GoalsEvent struct {
	ID                   string `json:"id"`                     // An ID that identifies this event.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // An ID that uniquely identifies the broadcaster.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s user handle.
	Type                 string `json:"type"`                   // The type of goal.
	Description          string `json:"description"`            // A description of the goal, if specified.
	IsAchieved           bool   `json:"is_achieved"`            // A Boolean value that indicates whether the broadcaster achieved their goal.
	CurrentAmount        int    `json:"current_amount"`         // The goal’s current value.
	TargetAmount         int    `json:"target_amount"`          // The goal’s target value.
	StartedAt            string `json:"started_at"`             // The UTC timestamp in RFC 3339 format, which indicates when the broadcaster created the goal.
	EndedAt              string `json:"ended_at"`               // The UTC timestamp in RFC 3339 format, which indicates when the broadcaster ended the goal.
}

type GoalsEventCondition struct{}

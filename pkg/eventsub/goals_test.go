package eventsub

import "testing"

func TestGoals(t *testing.T) {
	input := `{
        "id": "12345-cool-event",
        "broadcaster_user_id": "141981764",
        "broadcaster_user_name": "TwitchDev",
        "broadcaster_user_login": "twitchdev",
        "type": "subscription",
        "description": "Help me get partner!",
        "current_amount": 120,
        "target_amount": 220,
        "started_at": "2021-07-15T17:16:03.17106713Z"
    }`
	expected := GoalsEvent{
		ID:                   "12345-cool-event",
		BroadcasterUserID:    "141981764",
		BroadcasterUserName:  "TwitchDev",
		BroadcasterUserLogin: "twitchdev",
		Type:                 "subscription",
		Description:          "Help me get partner!",
		CurrentAmount:        120,
		TargetAmount:         220,
		StartedAt:            "2021-07-15T17:16:03.17106713Z",
	}

	validateInput(t, input, expected)
}

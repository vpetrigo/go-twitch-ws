package eventsub

import "testing"

func TestStreamOnline(t *testing.T) {
	input := `{
        "id": "9001",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "type": "live",
        "started_at": "2020-10-11T10:11:12.123Z"
    }`
	expected := StreamOnlineEvent{
		ID:                   "9001",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		Type:                 "live",
		StartedAt:            "2020-10-11T10:11:12.123Z",
	}

	validateInput(t, input, expected)
}

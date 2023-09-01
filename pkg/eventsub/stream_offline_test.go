package eventsub

import "testing"

func TestStreamOffline(t *testing.T) {
	input := `{
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User"
    }`
	expected := StreamOfflineEvent{
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
	}

	validateInput(t, input, expected)
}

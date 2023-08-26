package eventsub

import "testing"

func TestChannelGuestStarSessionEnd(t *testing.T) {
	input := `{
        "broadcaster_user_id": "1337",
        "broadcaster_user_name": "Cool_User",
        "broadcaster_user_login": "cool_user",
        "moderator_user_id": "1338",
        "moderator_user_name": "Cool_Mod",
        "moderator_user_login": "cool_mod",
        "session_id": "2KFRQbFtpmfyD3IevNRnCzOPRJI",
        "started_at": "2023-04-11T16:20:03.17106713Z",
        "ended_at": "2023-04-11T17:51:29.153485Z"
    }`
	expected := ChannelGuestStarSessionEndEvent{
		BroadcasterUserID:    "1337",
		BroadcasterUserName:  "Cool_User",
		BroadcasterUserLogin: "cool_user",
		SessionID:            "2KFRQbFtpmfyD3IevNRnCzOPRJI",
		StartedAt:            "2023-04-11T16:20:03.17106713Z",
		EndedAt:              "2023-04-11T17:51:29.153485Z",
	}

	validateInput(t, input, expected)
}

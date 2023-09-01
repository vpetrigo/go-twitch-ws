package eventsub

import "testing"

func TestChannelGuestStarGuestUpdate(t *testing.T) {
	input := `{
        "broadcaster_user_id": "1337",
        "broadcaster_user_name": "Cool_User",
        "broadcaster_user_login": "cool_user",
        "session_id": "2KFRQbFtpmfyD3IevNRnCzOPRJI",
        "moderator_user_id": "1312",
        "moderator_user_name": "Cool_Mod",
        "moderator_user_login": "cool_mod",
        "guest_user_id": "1234",
        "guest_user_name": "Cool_Guest",
        "guest_user_login": "cool_guest",
        "slot_id": "1",
        "state": "live",
        "host_video_enabled": true,
        "host_audio_enabled": true,
        "host_volume": 100
    }`
	expected := ChannelGuestStarGuestUpdateEvent{
		BroadcasterUserID:    "1337",
		BroadcasterUserName:  "Cool_User",
		BroadcasterUserLogin: "cool_user",
		SessionID:            "2KFRQbFtpmfyD3IevNRnCzOPRJI",
		ModeratorUserID:      "1312",
		ModeratorUserName:    "Cool_Mod",
		ModeratorUserLogin:   "cool_mod",
		GuestUserID:          "1234",
		GuestUserName:        "Cool_Guest",
		GuestUserLogin:       "cool_guest",
		SlotID:               "1",
		State:                "live",
		HostVideoEnabled:     true,
		HostAudioEnabled:     true,
		HostVolume:           100,
	}

	validateInput(t, input, expected)
}

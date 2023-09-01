package eventsub

import "testing"

func TestChannelGuestStarSettingsUpdate(t *testing.T) {
	input := `{
        "broadcaster_user_id": "1337",
        "broadcaster_user_name": "Cool_User",
        "broadcaster_user_login": "cool_user",
        "is_moderator_send_live_enabled": true,
        "slot_count": 5,
        "is_browser_source_audio_enabled": true,
        "group_layout": "tiled"
    }`
	expected := ChannelGuestStarSettingsUpdateEvent{
		BroadcasterUserID:           "1337",
		BroadcasterUserName:         "Cool_User",
		BroadcasterUserLogin:        "cool_user",
		IsModeratorSendLiveEnabled:  true,
		SlotCount:                   5,
		IsBrowserSourceAudioEnabled: true,
		GroupLayout:                 "tiled",
	}

	validateInput(t, input, expected)
}

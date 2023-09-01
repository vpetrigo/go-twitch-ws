package eventsub

import "testing"

func TestChannelModeratorAdd(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "mod_user",
        "user_name": "Mod_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User"
    }`
	expected := ChannelModeratorAddEvent{
		UserID:               "1234",
		UserLogin:            "mod_user",
		UserName:             "Mod_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
	}

	validateInput(t, input, expected)
}

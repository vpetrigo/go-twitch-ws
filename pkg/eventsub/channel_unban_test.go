package eventsub

import "testing"

func TestChannelUnban(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "cool_user",
        "user_name": "Cool_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User",
        "moderator_user_id": "1339",
        "moderator_user_login": "mod_user",
        "moderator_user_name": "Mod_User"
    }`
	expected := ChannelUnbanEvent{
		UserID:               "1234",
		UserLogin:            "cool_user",
		UserName:             "Cool_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
		ModeratorUserID:      "1339",
		ModeratorUserLogin:   "mod_user",
		ModeratorUserName:    "Mod_User",
	}

	validateInput(t, input, expected)
}

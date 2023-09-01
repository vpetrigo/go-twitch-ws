package eventsub

import "testing"

func TestChannelModeratorRemote(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "not_mod_user",
        "user_name": "Not_Mod_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User"
    }`
	expected := ChannelModeratorRemoveEvent{
		UserID:               "1234",
		UserLogin:            "not_mod_user",
		UserName:             "Not_Mod_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
	}

	validateInput(t, input, expected)
}

package eventsub

import (
	"testing"
)

func TestChannelBan(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "cool_user",
        "user_name": "Cool_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User",
        "moderator_user_id": "1339",
        "moderator_user_login": "mod_user",
        "moderator_user_name": "Mod_User",
        "reason": "Offensive language",
        "banned_at": "2020-07-15T18:15:11.17106713Z",
        "ends_at": "2020-07-15T18:16:11.17106713Z",
        "is_permanent": false
    }`
	expected := ChannelBanEvent{
		UserID:               "1234",
		UserLogin:            "cool_user",
		UserName:             "Cool_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
		ModeratorUserID:      "1339",
		ModeratorUserLogin:   "mod_user",
		ModeratorUserName:    "Mod_User",
		Reason:               "Offensive language",
		BannedAt:             "2020-07-15T18:15:11.17106713Z",
		EndsAt:               "2020-07-15T18:16:11.17106713Z",
		IsPermanent:          false,
	}

	validateInput(t, input, expected)
}

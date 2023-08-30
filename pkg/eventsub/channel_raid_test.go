package eventsub

import "testing"

func TestChannelRaid(t *testing.T) {
	input := `{
        "from_broadcaster_user_id": "1234",
        "from_broadcaster_user_login": "cool_user",
        "from_broadcaster_user_name": "Cool_User",
        "to_broadcaster_user_id": "1337",
        "to_broadcaster_user_login": "cooler_user",
        "to_broadcaster_user_name": "Cooler_User",
        "viewers": 9001
    }`
	expected := ChannelRaidEvent{
		FromBroadcasterUserID:    "1234",
		FromBroadcasterUserLogin: "cool_user",
		FromBroadcasterUserName:  "Cool_User",
		ToBroadcasterUserID:      "1337",
		ToBroadcasterUserLogin:   "cooler_user",
		ToBroadcasterUserName:    "Cooler_User",
		Viewers:                  9001,
	}

	validateInput(t, input, expected)
}

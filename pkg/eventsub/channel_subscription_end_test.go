package eventsub

import "testing"

func TestChannelSubscriptionEnd(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "cool_user",
        "user_name": "Cool_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User",
        "tier": "1000",
        "is_gift": false
    }`
	expected := ChannelSubscriptionEndEvent{
		UserID:               "1234",
		UserLogin:            "cool_user",
		UserName:             "Cool_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
		Tier:                 "1000",
		IsGift:               false,
	}

	validateInput(t, input, expected)
}

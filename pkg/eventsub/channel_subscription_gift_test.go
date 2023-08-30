package eventsub

import "testing"

func TestChannelSubscriptionGift(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "cool_user",
        "user_name": "Cool_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User",
        "total": 2,
        "tier": "1000",
        "cumulative_total": 284,
        "is_anonymous": false
    }`
	expected := ChannelSubscriptionGiftEvent{
		UserID:               "1234",
		UserLogin:            "cool_user",
		UserName:             "Cool_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
		Total:                2,
		Tier:                 "1000",
		CumulativeTotal:      284,
		IsAnonymous:          false,
	}

	validateInput(t, input, expected)
}

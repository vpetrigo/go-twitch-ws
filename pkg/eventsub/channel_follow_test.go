package eventsub

import (
	"testing"
)

func TestChannelFollow(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "cool_user",
        "user_name": "Cool_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User",
        "followed_at": "2020-07-15T18:16:11.17106713Z"
    }`
	expected := ChannelFollowEvent{
		UserID:               "1234",
		UserLogin:            "cool_user",
		UserName:             "Cool_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
		FollowedAt:           "2020-07-15T18:16:11.17106713Z",
	}

	validateInput(t, input, expected)
}

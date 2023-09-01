package eventsub

import (
	"testing"
)

func TestChannelCheer(t *testing.T) {
	input := `{
        "is_anonymous": false,
        "user_id": "1234",          
        "user_login": "cool_user",  
        "user_name": "Cool_User",   
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User",
        "message": "pogchamp",
        "bits": 1000
    }`
	expected := ChannelCheerEvent{
		IsAnonymous:          false,
		UserID:               "1234",
		UserLogin:            "cool_user",
		UserName:             "Cool_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
		Message:              "pogchamp",
		Bits:                 1000,
	}

	validateInput(t, input, expected)
}

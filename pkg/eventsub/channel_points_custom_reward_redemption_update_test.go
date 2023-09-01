package eventsub

import "testing"

func TestChannelPointsCustomRewardRedemptionUpdate(t *testing.T) {
	input := `{
        "id": "17fa2df1-ad76-4804-bfa5-a40ef63efe63",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "user_id": "9001",
        "user_login": "cooler_user",
        "user_name": "Cooler_User",
        "user_input": "pogchamp",
        "status": "fulfilled",
        "reward": {
            "id": "92af127c-7326-4483-a52b-b0da0be61c01",
            "title": "title",
            "cost": 100,
            "prompt": "reward prompt"
        },
        "redeemed_at": "2020-07-15T17:16:03.17106713Z"
    }`
	expected := ChannelPointsCustomRewardRedemptionUpdateEvent{
		ID:                   "17fa2df1-ad76-4804-bfa5-a40ef63efe63",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		UserID:               "9001",
		UserLogin:            "cooler_user",
		UserName:             "Cooler_User",
		UserInput:            "pogchamp",
		Status:               "fulfilled",
		Reward: Reward{
			ID:     "92af127c-7326-4483-a52b-b0da0be61c01",
			Title:  "title",
			Cost:   100,
			Prompt: "reward prompt",
		},
		RedeemedAt: "2020-07-15T17:16:03.17106713Z",
	}

	validateInput(t, input, expected)
}

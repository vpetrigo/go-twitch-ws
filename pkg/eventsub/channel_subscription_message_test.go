package eventsub

import (
	"encoding/json"
	"testing"
)

// TODO: Update Message to have array of emotes.
func TestChannelSubscriptionMessage(t *testing.T) {
	input := `{
        "user_id": "1234",
        "user_login": "cool_user",
        "user_name": "Cool_User",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cooler_user",
        "broadcaster_user_name": "Cooler_User",
        "tier": "1000",
        "message": {
            "text": "Love the stream! FevziGG",
            "emotes": [
                {
                    "begin": 23,
                    "end": 30,
                    "id": "302976485"
                }
            ]
        },
        "cumulative_months": 15,
        "streak_months": 1,
        "duration_months": 6
    }`
	expected := ChannelSubscriptionMessageEvent{
		UserID:               "1234",
		UserLogin:            "cool_user",
		UserName:             "Cool_User",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cooler_user",
		BroadcasterUserName:  "Cooler_User",
		Tier:                 "1000",
		Message: Message{
			Text: "Love the stream! FevziGG",
			Emotes: []Emotes{
				{
					Begin: 23,
					End:   30,
					ID:    "302976485",
				},
			},
		},
		CumulativeMonths: 15,
		StreakMonths:     1,
		DurationMonths:   6,
	}
	var actual ChannelSubscriptionMessageEvent
	err := json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	userIDEq := expected.UserID == actual.UserID
	userLoginEq := expected.UserLogin == actual.UserLogin
	userNameEq := expected.UserName == actual.UserName
	broadcasterUserIDEq := expected.BroadcasterUserID == actual.BroadcasterUserID
	broadcasterUserLoginEq := expected.BroadcasterUserLogin == actual.BroadcasterUserLogin
	tierEq := expected.Tier == actual.Tier
	messageTextEq := expected.Message.Text == actual.Message.Text
	messageEmotesEq := len(expected.Message.Emotes) == len(actual.Message.Emotes)

	if messageEmotesEq {
		for i, v := range expected.Message.Emotes {
			if v != actual.Message.Emotes[i] {
				messageEmotesEq = false
				break
			}
		}
	}

	cumulativeMonthsEq := expected.CumulativeMonths == actual.CumulativeMonths
	streakMonthsEq := expected.StreakMonths == actual.StreakMonths
	durationMonthsEq := expected.DurationMonths == actual.DurationMonths

	if !(userIDEq &&
		userLoginEq &&
		userNameEq &&
		broadcasterUserIDEq &&
		broadcasterUserLoginEq &&
		tierEq &&
		messageTextEq &&
		messageEmotesEq &&
		cumulativeMonthsEq &&
		streakMonthsEq &&
		durationMonthsEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

package eventsub

import (
	"encoding/json"
	"testing"
)

func TestChannelPredictionLock(t *testing.T) {
	input := `{
        "id": "1243456",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "title": "Aren’t shoes just really hard socks?",
        "outcomes": [
            {
                "id": "1243456",
                "title": "Yeah!",
                "color": "blue",
                "users": 10,
                "channel_points": 15000,
                "top_predictors": [
                    {
                        "user_name": "Cool_User",
                        "user_login": "cool_user",
                        "user_id": "1234",
                        "channel_points_won": null,
                        "channel_points_used": 500
                    },
                    {
                        "user_name": "Coolest_User",
                        "user_login": "coolest_user",
                        "user_id": "1236",
                        "channel_points_won": null,
                        "channel_points_used": 200
                    }
                ]
            },
            {
                "id": "2243456",
                "title": "No!",
                "color": "pink",
                "top_predictors": [
                    {
                        "user_name": "Cooler_User",
                        "user_login": "cooler_user",
                        "user_id": "12345",
                        "channel_points_won": null,
                        "channel_points_used": 5000
                    }
                ]
            }
        ],
        "started_at": "2020-07-15T17:16:03.17106713Z",
        "locked_at": "2020-07-15T17:21:03.17106713Z"
    }`
	expected := ChannelPredictionLockEvent{
		ID:                   "1243456",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		Title:                "Aren’t shoes just really hard socks?",
		Outcomes: []Outcomes{
			{
				ID:            "1243456",
				Title:         "Yeah!",
				Color:         "blue",
				Users:         10,
				ChannelPoints: 15000,
				TopPredictors: []TopPredictors{
					{
						UserName:          "Cool_User",
						UserLogin:         "cool_user",
						UserID:            "1234",
						ChannelPointsWon:  0,
						ChannelPointsUsed: 500,
					},
					{
						UserName:          "Coolest_User",
						UserLogin:         "coolest_user",
						UserID:            "1236",
						ChannelPointsWon:  0,
						ChannelPointsUsed: 200,
					},
				},
			},
			{
				ID:            "2243456",
				Title:         "No!",
				Users:         0,
				ChannelPoints: 0,
				Color:         "pink",
				TopPredictors: []TopPredictors{
					{
						UserName:          "Cooler_User",
						UserLogin:         "cooler_user",
						UserID:            "12345",
						ChannelPointsWon:  0,
						ChannelPointsUsed: 5000,
					},
				},
			},
		},
		StartedAt: "2020-07-15T17:16:03.17106713Z",
		LockedAt:  "2020-07-15T17:21:03.17106713Z",
	}
	var actual ChannelPredictionLockEvent
	err := json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	idEq := expected.ID == actual.ID
	broadcasterUserIDEq := expected.BroadcasterUserID == actual.BroadcasterUserID
	broadcasterUserLoginEq := expected.BroadcasterUserLogin == actual.BroadcasterUserLogin
	broadcasterUserNameEq := expected.BroadcasterUserName == actual.BroadcasterUserName
	titleEq := expected.Title == actual.Title

	outcomesEq := len(expected.Outcomes) == len(actual.Outcomes)

	if outcomesEq {
	outer:
		for i, v := range expected.Outcomes {
			if !(v.ID == actual.Outcomes[i].ID && v.Title == actual.Outcomes[i].Title &&
				v.ChannelPoints == actual.Outcomes[i].ChannelPoints && v.Color == actual.Outcomes[i].Color &&
				v.Users == actual.Outcomes[i].Users &&
				len(v.TopPredictors) == len(actual.Outcomes[i].TopPredictors)) {
				outcomesEq = false
				break
			}

			for j, p := range v.TopPredictors {
				if p != actual.Outcomes[i].TopPredictors[j] {
					outcomesEq = false
					break outer
				}
			}
		}
	}

	startedAtEq := expected.StartedAt == actual.StartedAt
	lockedAtEq := expected.LockedAt == actual.LockedAt

	if !(idEq &&
		broadcasterUserIDEq &&
		broadcasterUserLoginEq &&
		broadcasterUserNameEq &&
		titleEq &&
		outcomesEq &&
		startedAtEq &&
		lockedAtEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

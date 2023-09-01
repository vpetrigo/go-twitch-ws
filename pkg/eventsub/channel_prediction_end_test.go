package eventsub

import (
	"encoding/json"
	"testing"
)

func TestChannelPredictionEnd(t *testing.T) {
	input := `{
	"id": "1243456",
	"broadcaster_user_id": "1337",
	"broadcaster_user_login": "cool_user",
	"broadcaster_user_name": "Cool_User",
	"title": "Aren’t shoes just really hard socks?",
	"winning_outcome_id": "12345",
	"outcomes": [{
			"id": "12345",
			"title": "Yeah!",
			"color": "blue",
			"users": 2,
			"channel_points": 15000,
			"top_predictors": [{
					"user_name": "Cool_User",
					"user_login": "cool_user",
					"user_id": "1234",
					"channel_points_won": 10000,
					"channel_points_used": 500
				},
				{
					"user_name": "Coolest_User",
					"user_login": "coolest_user",
					"user_id": "1236",
					"channel_points_won": 5000,
					"channel_points_used": 100
				}
			]
		},
		{
			"id": "22435",
			"title": "No!",
			"users": 2,
			"channel_points": 200,
			"color": "pink",
			"top_predictors": [{
					"user_name": "Cooler_User",
					"user_login": "cooler_user",
					"user_id": "12345",
					"channel_points_won": null,
					"channel_points_used": 100
				},
				{
					"user_name": "Elite_User",
					"user_login": "elite_user",
					"user_id": "1337",
					"channel_points_won": null,
					"channel_points_used": 100
				}
			]
		}
	],
	"status": "resolved",
	"started_at": "2020-07-15T17:16:03.17106713Z",
	"ended_at": "2020-07-15T17:16:11.17106713Z"
}`
	expected := ChannelPredictionEndEvent{
		ID:                   "1243456",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		Title:                "Aren’t shoes just really hard socks?",
		WinningOutcomeID:     "12345",
		Outcomes: []Outcomes{
			{
				ID:            "12345",
				Title:         "Yeah!",
				Color:         "blue",
				Users:         2,
				ChannelPoints: 15000,
				TopPredictors: []TopPredictors{
					{
						UserName:          "Cool_User",
						UserLogin:         "cool_user",
						UserID:            "1234",
						ChannelPointsWon:  10000,
						ChannelPointsUsed: 500,
					},
					{
						UserName:          "Coolest_User",
						UserLogin:         "coolest_user",
						UserID:            "1236",
						ChannelPointsWon:  5000,
						ChannelPointsUsed: 100,
					},
				},
			},
			{
				ID:            "22435",
				Title:         "No!",
				Users:         2,
				ChannelPoints: 200,
				Color:         "pink",
				TopPredictors: []TopPredictors{
					{
						UserName:          "Cooler_User",
						UserLogin:         "cooler_user",
						UserID:            "12345",
						ChannelPointsWon:  0,
						ChannelPointsUsed: 100,
					},
					{
						UserName:          "Elite_User",
						UserLogin:         "elite_user",
						UserID:            "1337",
						ChannelPointsWon:  0,
						ChannelPointsUsed: 100,
					},
				},
			},
		},
		Status:    "resolved",
		StartedAt: "2020-07-15T17:16:03.17106713Z",
		EndedAt:   "2020-07-15T17:16:11.17106713Z",
	}
	var actual ChannelPredictionEndEvent
	err := json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	idEq := expected.ID == actual.ID
	broadcasterUserIDEq := expected.BroadcasterUserID == actual.BroadcasterUserID
	broadcasterUserLoginEq := expected.BroadcasterUserLogin == actual.BroadcasterUserLogin
	broadcasterUserNameEq := expected.BroadcasterUserName == actual.BroadcasterUserName
	titleEq := expected.Title == actual.Title
	winningOutcomeIDEq := expected.WinningOutcomeID == actual.WinningOutcomeID

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

	statusEq := expected.Status == actual.Status
	startedAtEq := expected.StartedAt == actual.StartedAt
	endedAtEq := expected.EndedAt == actual.EndedAt

	if !(idEq &&
		broadcasterUserIDEq &&
		broadcasterUserLoginEq &&
		broadcasterUserNameEq &&
		titleEq &&
		winningOutcomeIDEq &&
		outcomesEq &&
		statusEq &&
		startedAtEq &&
		endedAtEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

package eventsub

import (
	"encoding/json"
	"testing"
)

func TestChannelPredictionBegin(t *testing.T) {
	input := `{
        "id": "1243456",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "title": "Aren’t shoes just really hard socks?",
        "outcomes": [
            {"id": "1243456", "title": "Yeah!", "color": "blue"},
            {"id": "2243456", "title": "No!", "color": "pink"}
        ],
        "started_at": "2020-07-15T17:16:03.17106713Z",
        "locks_at": "2020-07-15T17:21:03.17106713Z"
    }`
	expected := ChannelPredictionBeginEvent{
		ID:                   "1243456",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		Title:                "Aren’t shoes just really hard socks?",
		Outcomes: []Outcomes{
			{
				ID:    "1243456",
				Title: "Yeah!",
				Color: "blue",
			},
			{
				ID:    "2243456",
				Title: "No!",
				Color: "pink",
			},
		},
		StartedAt: "2020-07-15T17:16:03.17106713Z",
		LocksAt:   "2020-07-15T17:21:03.17106713Z",
	}
	var actual ChannelPredictionBeginEvent
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
	locksAtEq := expected.LocksAt == actual.LocksAt

	if !(idEq &&
		broadcasterUserIDEq &&
		broadcasterUserLoginEq &&
		broadcasterUserNameEq &&
		titleEq &&
		outcomesEq &&
		startedAtEq &&
		locksAtEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

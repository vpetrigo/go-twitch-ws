package eventsub

import (
	"encoding/json"
	"testing"
)

func TestHypeTrainEnd(t *testing.T) {
	input := `{
        "id": "1b0AsbInCHZW2SQFQkCzqN07Ib2",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "level": 2,
        "total": 137,
        "top_contributions": [
            { "user_id": "123", "user_login": "pogchamp", "user_name": "PogChamp", "type": "bits", "total": 50 },
            { "user_id": "456", "user_login": "kappa", "user_name": "Kappa", "type": "subscription", "total": 45 }
        ],
        "started_at": "2020-07-15T17:16:03.17106713Z",
        "ended_at": "2020-07-15T17:16:11.17106713Z",
        "cooldown_ends_at": "2020-07-15T18:16:11.17106713Z"
    }`
	expected := HypeTrainEndEvent{
		ID:                   "1b0AsbInCHZW2SQFQkCzqN07Ib2",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		Level:                2,
		Total:                137,
		TopContributions: []TopContributions{
			{UserID: "123", UserLogin: "pogchamp", UserName: "PogChamp", Type: "bits", Total: 50},
			{UserID: "456", UserLogin: "kappa", UserName: "Kappa", Type: "subscription", Total: 45},
		},
		StartedAt:      "2020-07-15T17:16:03.17106713Z",
		EndedAt:        "2020-07-15T17:16:11.17106713Z",
		CooldownEndsAt: "2020-07-15T18:16:11.17106713Z",
	}
	var actual HypeTrainEndEvent

	err := json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	idEq := actual.ID == expected.ID
	broadcasterIDEq := actual.BroadcasterUserID == expected.BroadcasterUserID
	broadcasterUserLoginEq := actual.BroadcasterUserLogin == expected.BroadcasterUserLogin
	broadcasterUserNameEq := actual.BroadcasterUserName == expected.BroadcasterUserName
	levelEq := actual.Level == expected.Level
	totalEq := actual.Total == expected.Total
	topContributionsEq := len(actual.TopContributions) == len(expected.TopContributions)

	if topContributionsEq {
		for i, v := range expected.TopContributions {
			if v != actual.TopContributions[i] {
				topContributionsEq = false
				break
			}
		}
	}

	startedAtEq := actual.StartedAt == expected.StartedAt
	expiresAtEq := actual.EndedAt == expected.EndedAt
	cooldownEq := actual.CooldownEndsAt == expected.CooldownEndsAt

	if !(idEq &&
		broadcasterIDEq &&
		broadcasterUserLoginEq &&
		broadcasterUserNameEq &&
		levelEq &&
		totalEq &&
		topContributionsEq &&
		startedAtEq &&
		expiresAtEq &&
		cooldownEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

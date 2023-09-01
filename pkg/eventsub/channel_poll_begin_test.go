package eventsub

import (
	"encoding/json"
	"testing"
)

func TestChannelPoolBegin(t *testing.T) {
	input := `{
        "id": "1243456",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "title": "Aren’t shoes just really hard socks?",
        "choices": [
            {"id": "123", "title": "Yeah!"},
            {"id": "124", "title": "No!"},
            {"id": "125", "title": "Maybe!"}
        ],
        "bits_voting": {
            "is_enabled": true,
            "amount_per_vote": 10
        },
        "channel_points_voting": {
            "is_enabled": true,
            "amount_per_vote": 10
        },
        "started_at": "2020-07-15T17:16:03.17106713Z",
        "ends_at": "2020-07-15T17:16:08.17106713Z"
    }`
	expected := ChannelPollBeginEvent{
		ID:                   "1243456",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		Title:                "Aren’t shoes just really hard socks?",
		BitsVoting: BitsVoting{
			IsEnabled:     true,
			AmountPerVote: 10,
		},
		Choices: []Choices{
			{
				ID:    "123",
				Title: "Yeah!",
			},
			{
				ID:    "124",
				Title: "No!",
			},
			{
				ID:    "125",
				Title: "Maybe!",
			},
		},
		ChannelPointsVoting: ChannelPointsVoting{
			IsEnabled:     true,
			AmountPerVote: 10,
		},
		StartedAt: "2020-07-15T17:16:03.17106713Z",
		EndsAt:    "2020-07-15T17:16:08.17106713Z",
	}
	var actual ChannelPollBeginEvent
	err := json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	idEq := expected.ID == actual.ID
	broadcasterUserIDEq := expected.BroadcasterUserID == actual.BroadcasterUserID
	broadcasterUserNameEq := expected.BroadcasterUserName == actual.BroadcasterUserName
	broadcasterUserLoginEq := expected.BroadcasterUserLogin == actual.BroadcasterUserLogin
	titleEq := expected.Title == actual.Title
	bitsVotingEq := expected.BitsVoting == actual.BitsVoting
	choicesEq := len(expected.Choices) == len(actual.Choices)

	if choicesEq {
		for i, v := range expected.Choices {
			if v != actual.Choices[i] {
				choicesEq = false
				break
			}
		}
	}

	startedAt := expected.StartedAt == actual.StartedAt
	endsAtEq := expected.EndsAt == actual.EndsAt

	if !(idEq &&
		broadcasterUserIDEq &&
		broadcasterUserNameEq &&
		broadcasterUserLoginEq &&
		titleEq &&
		bitsVotingEq &&
		choicesEq &&
		startedAt &&
		endsAtEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

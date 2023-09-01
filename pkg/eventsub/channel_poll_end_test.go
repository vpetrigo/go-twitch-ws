package eventsub

import (
	"encoding/json"
	"testing"
)

func TestChannelPoolEnd(t *testing.T) {
	input := `{
        "id": "1243456",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "title": "Aren’t shoes just really hard socks?",
        "choices": [
            {"id": "123", "title": "Blue", "bits_votes": 50, "channel_points_votes": 70, "votes": 120},
            {"id": "124", "title": "Yellow", "bits_votes": 100, "channel_points_votes": 40, "votes": 140},
            {"id": "125", "title": "Green", "bits_votes": 10, "channel_points_votes": 70, "votes": 80}
        ],
        "bits_voting": {
            "is_enabled": true,
            "amount_per_vote": 10
        },
        "channel_points_voting": {
            "is_enabled": true,
            "amount_per_vote": 10
        },
        "status": "completed",
        "started_at": "2020-07-15T17:16:03.17106713Z",
        "ended_at": "2020-07-15T17:16:11.17106713Z"
    }`
	expected := ChannelPollEndEvent{
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
				ID:                 "123",
				Title:              "Blue",
				BitsVotes:          50,
				ChannelPointsVotes: 70,
				Votes:              120,
			},
			{
				ID:                 "124",
				Title:              "Yellow",
				BitsVotes:          100,
				ChannelPointsVotes: 40,
				Votes:              140,
			},
			{
				ID:                 "125",
				Title:              "Green",
				BitsVotes:          10,
				ChannelPointsVotes: 70,
				Votes:              80,
			},
		},
		ChannelPointsVoting: ChannelPointsVoting{
			IsEnabled:     true,
			AmountPerVote: 10,
		},
		Status:    "completed",
		StartedAt: "2020-07-15T17:16:03.17106713Z",
		EndedAt:   "2020-07-15T17:16:11.17106713Z",
	}
	var actual ChannelPollEndEvent
	err := json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	idEq := expected.ID == actual.ID
	broadcasterUserIdEq := expected.BroadcasterUserID == actual.BroadcasterUserID
	broadcasterUserNameEq := expected.BroadcasterUserName == actual.BroadcasterUserName
	broadcasterUserLoginEq := expected.BroadcasterUserLogin == actual.BroadcasterUserLogin
	titleEq := expected.Title == actual.Title
	bitsVotingEq := expected.BitsVoting == actual.BitsVoting
	choicesEq := len(expected.Choices) == len(expected.Choices)

	if choicesEq {
		for i, v := range expected.Choices {
			if v != actual.Choices[i] {
				choicesEq = false
				break
			}
		}
	}

	statusEq := expected.Status == actual.Status
	startedAt := expected.StartedAt == actual.StartedAt
	endedAtEq := expected.EndedAt == actual.EndedAt

	if !(idEq &&
		broadcasterUserIdEq &&
		broadcasterUserNameEq &&
		broadcasterUserLoginEq &&
		titleEq &&
		bitsVotingEq &&
		choicesEq &&
		statusEq &&
		startedAt &&
		endedAtEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

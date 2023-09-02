package eventsub

import "testing"

func TestShieldMode(t *testing.T) {
	input := []string{
		`{
    "broadcaster_user_id": "12345",
    "broadcaster_user_name": "SimplySimple",
    "broadcaster_user_login": "simplysimple",
    "moderator_user_id": "98765",
    "moderator_user_name": "ParticularlyParticular123",
    "moderator_user_login": "particularlyparticular123",
    "started_at": "2022-07-26T17:00:03.17106713Z"
  }`,
		`{
    "broadcaster_user_id": "12345",
    "broadcaster_user_name": "SimplySimple",
    "broadcaster_user_login": "simplysimple",
    "moderator_user_id": "98765",
    "moderator_user_name": "ParticularlyParticular123",
    "moderator_user_login": "particularlyparticular123",
    "ended_at": "2022-07-27T01:30:23.17106713Z"
  }`,
	}
	expected := []ShieldModeEvent{
		{
			BroadcasterUserID:    "12345",
			BroadcasterUserName:  "SimplySimple",
			BroadcasterUserLogin: "simplysimple",
			ModeratorUserID:      "98765",
			ModeratorUserName:    "ParticularlyParticular123",
			ModeratorUserLogin:   "particularlyparticular123",
			StartedAt:            "2022-07-26T17:00:03.17106713Z",
		},
		{
			BroadcasterUserID:    "12345",
			BroadcasterUserName:  "SimplySimple",
			BroadcasterUserLogin: "simplysimple",
			ModeratorUserID:      "98765",
			ModeratorUserName:    "ParticularlyParticular123",
			ModeratorUserLogin:   "particularlyparticular123",
			EndedAt:              "2022-07-27T01:30:23.17106713Z",
		},
	}

	for i, v := range expected {
		validateInput(t, input[i], v)
	}
}

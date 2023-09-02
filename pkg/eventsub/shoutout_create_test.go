package eventsub

import "testing"

func TestShoutoutCreate(t *testing.T) {
	input := `{
    "broadcaster_user_id": "12345",
    "broadcaster_user_name": "SimplySimple",
    "broadcaster_user_login": "simplysimple",
    "moderator_user_id": "98765",
    "moderator_user_name": "ParticularlyParticular123",
    "moderator_user_login": "particularlyparticular123",
    "to_broadcaster_user_id": "626262",
    "to_broadcaster_user_name": "SandySanderman",
    "to_broadcaster_user_login": "sandysanderman",
    "started_at": "2022-07-26T17:00:03.17106713Z",
    "viewer_count": 860,
    "cooldown_ends_at": "2022-07-26T17:02:03.17106713Z",
    "target_cooldown_ends_at":"2022-07-26T18:00:03.17106713Z"
  }`
	expected := ShoutoutCreateEvent{
		BroadcasterUserID:      "12345",
		BroadcasterUserName:    "SimplySimple",
		BroadcasterUserLogin:   "simplysimple",
		ModeratorUserID:        "98765",
		ModeratorUserName:      "ParticularlyParticular123",
		ModeratorUserLogin:     "particularlyparticular123",
		ToBroadcasterUserID:    "626262",
		ToBroadcasterUserName:  "SandySanderman",
		ToBroadcasterUserLogin: "sandysanderman",
		StartedAt:              "2022-07-26T17:00:03.17106713Z",
		ViewerCount:            860,
		CooldownEndsAt:         "2022-07-26T17:02:03.17106713Z",
		TargetCooldownEndsAt:   "2022-07-26T18:00:03.17106713Z",
	}

	validateInput(t, input, expected)
}

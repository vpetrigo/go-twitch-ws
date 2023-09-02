package eventsub

import "testing"

func TestShoutoutReceived(t *testing.T) {
	input := `{
    "broadcaster_user_id": "626262",
    "broadcaster_user_name": "SandySanderman",
    "broadcaster_user_login": "sandysanderman",
    "from_broadcaster_user_id": "12345",
    "from_broadcaster_user_name": "SimplySimple",
    "from_broadcaster_user_login": "simplysimple",
    "viewer_count": 860,
    "started_at": "2022-07-26T17:00:03.17106713Z"
  }`
	expected := ShoutoutReceivedEvent{
		BroadcasterUserID:        "626262",
		BroadcasterUserName:      "SandySanderman",
		BroadcasterUserLogin:     "sandysanderman",
		FromBroadcasterUserID:    "12345",
		FromBroadcasterUserName:  "SimplySimple",
		FromBroadcasterUserLogin: "simplysimple",
		ViewerCount:              860,
		StartedAt:                "2022-07-26T17:00:03.17106713Z",
	}

	validateInput(t, input, expected)
}

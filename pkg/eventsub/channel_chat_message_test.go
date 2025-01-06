package eventsub

import (
	"encoding/json"
	"testing"
)

func TestChannelChatMessage(t *testing.T) {
	input := `
	{
	"broadcaster_user_id": "1971641",
	"broadcaster_user_login": "streamer",
	"broadcaster_user_name": "streamer",
	"chatter_user_id": "4145994",
	"chatter_user_login": "viewer32",
	"chatter_user_name": "viewer32",
	"message_id": "cc106a89-1814-919d-454c-f4f2f970aae7",
	"message": {
	  "text": "Hi chat",
	  "fragments": [
		{
		  "type": "text",
		  "text": "Hi chat",
		  "cheermote": null,
		  "emote": null,
		  "mention": null
		}
	  ]
	},
	"color": "#00FF7F",
	"badges": [
	  {
		"set_id": "moderator",
		"id": "1",
		"info": ""
	  },
	  {
		"set_id": "subscriber",
		"id": "12",
		"info": "16"
	  },
	  {
		"set_id": "sub-gifter",
		"id": "1",
		"info": ""
	  }
	],
	"message_type": "text",
	"cheer": null,
	"reply": null,
	"channel_points_custom_reward_id": null,
	"source_broadcaster_user_id": null,
	"source_broadcaster_user_login": null,
	"source_broadcaster_user_name": null,
	"source_message_id": null,
	"source_badges": null
	}`

	expected := ChannelChatMessage{
		BroadcasterUserID:    "1971641",
		BroadcasterUserName:  "streamer",
		BroadcasterUserLogin: "streamer",
		ChatterUserID:        "4145994",
		ChatterUserLogin:     "viewer32",
		ChatterUserName:      "viewer32",
		MessageID:            "cc106a89-1814-919d-454c-f4f2f970aae7",
		Message: ChatMessage{
			Text: "Hi chat",
			Fragments: []MessageFragment{
				{
					Type: "text", Text: "Hi chat", Cheermote: nil, Emote: nil, Mention: nil,
				},
			},
		},
		Color: "#00FF7F",
		Badges: []ChatBadge{
			{
				SetID: "moderator",
				ID:    "1",
				Info:  "",
			},
			{
				SetID: "subscriber",
				ID:    "12",
				Info:  "16",
			},
			{
				SetID: "sub-gifter",
				ID:    "1",
				Info:  "",
			},
		},
		MessageType:                 "text",
		Cheer:                       nil,
		Reply:                       nil,
		ChannelPointsCustomRewardID: nil,
		SourceBroadcasterUserID:     nil,
		SourceBroadcasterUserLogin:  nil,
		SourceBroadcasterUserName:   nil,
		SourceMessageID:             nil,
		SourceBadges:                nil,
	}

	var inputEvent ChannelChatMessage
	err := json.Unmarshal([]byte(input), &inputEvent)

	if err != nil {
		t.Fatal(err)
	}

	// Comparing top-level fields and saving results
	broadcasterUserIDEq := inputEvent.BroadcasterUserID == expected.BroadcasterUserID
	broadcasterUserLoginEq := inputEvent.BroadcasterUserLogin == expected.BroadcasterUserLogin
	broadcasterUserNameEq := inputEvent.BroadcasterUserName == expected.BroadcasterUserName
	chatterUserIDEq := inputEvent.ChatterUserID == expected.ChatterUserID
	chatterUserLoginEq := inputEvent.ChatterUserLogin == expected.ChatterUserLogin
	chatterUserNameEq := inputEvent.ChatterUserName == expected.ChatterUserName
	messageIDEq := inputEvent.MessageID == expected.MessageID
	messageTextEq := inputEvent.Message.Text == expected.Message.Text
	colorEq := inputEvent.Color == expected.Color
	messageTypeEq := inputEvent.MessageType == expected.MessageType
	cheerEq := inputEvent.Cheer == expected.Cheer
	replyEq := inputEvent.Reply == expected.Reply
	channelPointsRewardIDEq := inputEvent.ChannelPointsCustomRewardID == expected.ChannelPointsCustomRewardID
	sourceBroadcasterUserIDEq := inputEvent.SourceBroadcasterUserID == expected.SourceBroadcasterUserID
	sourceBroadcasterUserLoginEq := inputEvent.SourceBroadcasterUserLogin == expected.SourceBroadcasterUserLogin
	sourceBroadcasterUserNameEq := inputEvent.SourceBroadcasterUserName == expected.SourceBroadcasterUserName
	sourceMessageIDEq := inputEvent.SourceMessageID == expected.SourceMessageID

	// Check SourceBadges comparison
	sourceBadgesEq := (inputEvent.SourceBadges == nil && expected.SourceBadges == nil) ||
		(inputEvent.SourceBadges != nil && expected.SourceBadges != nil && *inputEvent.SourceBadges == *expected.SourceBadges)

	// Comparing Message.Fragments
	messageFragmentsEq := len(inputEvent.Message.Fragments) == len(expected.Message.Fragments)
	if messageFragmentsEq {
		for i := range inputEvent.Message.Fragments {
			if inputEvent.Message.Fragments[i] != expected.Message.Fragments[i] {
				messageFragmentsEq = false
				break
			}
		}
	}

	// Comparing Badges
	badgesEq := len(inputEvent.Badges) == len(expected.Badges)
	if badgesEq {
		for i := range inputEvent.Badges {
			if inputEvent.Badges[i] != expected.Badges[i] {
				badgesEq = false
				break
			}
		}
	}

	// Combine all checks
	if !(broadcasterUserIDEq &&
		broadcasterUserLoginEq &&
		broadcasterUserNameEq &&
		chatterUserIDEq &&
		chatterUserLoginEq &&
		chatterUserNameEq &&
		messageIDEq &&
		messageTextEq &&
		messageFragmentsEq &&
		colorEq &&
		badgesEq &&
		messageTypeEq &&
		cheerEq &&
		replyEq &&
		channelPointsRewardIDEq &&
		sourceBroadcasterUserIDEq &&
		sourceBroadcasterUserLoginEq &&
		sourceBroadcasterUserNameEq &&
		sourceMessageIDEq &&
		sourceBadgesEq) {
		t.Fatal(eventMismatchErrorMessage(inputEvent, expected))
	}
}

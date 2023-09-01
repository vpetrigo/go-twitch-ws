package eventsub

import "testing"

func TestChannelPointsCustomRewardUpdate(t *testing.T) {
	input := `{
        "id": "9001",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "is_enabled": true,
        "is_paused": false,
        "is_in_stock": true,
        "title": "Cool Reward",
        "cost": 100,
        "prompt": "reward prompt",
        "is_user_input_required": true,
        "should_redemptions_skip_request_queue": false,
        "cooldown_expires_at": "2019-11-16T10:11:12.634234626Z",
        "redemptions_redeemed_current_stream": 123,
        "max_per_stream": {
            "is_enabled": true,
            "value": 1000
        },
        "max_per_user_per_stream": {
            "is_enabled": true,
            "value": 1000
        },
        "global_cooldown": {
            "is_enabled": true,
            "seconds": 1000
        },
        "background_color": "#FA1ED2",
        "image": {
            "url_1x": "https://static-cdn.jtvnw.net/image-1.png",
            "url_2x": "https://static-cdn.jtvnw.net/image-2.png",
            "url_4x": "https://static-cdn.jtvnw.net/image-4.png"
        },
        "default_image": {
            "url_1x": "https://static-cdn.jtvnw.net/default-1.png",
            "url_2x": "https://static-cdn.jtvnw.net/default-2.png",
            "url_4x": "https://static-cdn.jtvnw.net/default-4.png"
        }
    }`
	expected := ChannelPointsCustomRewardUpdateEvent{
		ID:                                "9001",
		BroadcasterUserID:                 "1337",
		BroadcasterUserLogin:              "cool_user",
		BroadcasterUserName:               "Cool_User",
		IsEnabled:                         true,
		IsPaused:                          false,
		IsInStock:                         true,
		Title:                             "Cool Reward",
		Cost:                              100,
		Prompt:                            "reward prompt",
		IsUserInputRequired:               true,
		ShouldRedemptionsSkipRequestQueue: false,
		CooldownExpiresAt:                 "2019-11-16T10:11:12.634234626Z",
		RedemptionsRedeemedCurrentStream:  123,
		MaxPerStream: MaxPerStream{
			IsEnabled: true,
			Value:     1000,
		},
		MaxPerUserPerStream: MaxPerUserPerStream{
			IsEnabled: true,
			Value:     1000,
		},
		GlobalCooldown: GlobalCooldown{
			IsEnabled: true,
			Seconds:   1000,
		},
		BackgroundColor: "#FA1ED2",
		Image: Image{
			URL1X: "https://static-cdn.jtvnw.net/image-1.png",
			URL2X: "https://static-cdn.jtvnw.net/image-2.png",
			URL4X: "https://static-cdn.jtvnw.net/image-4.png",
		},
		DefaultImage: Image{
			URL1X: "https://static-cdn.jtvnw.net/default-1.png",
			URL2X: "https://static-cdn.jtvnw.net/default-2.png",
			URL4X: "https://static-cdn.jtvnw.net/default-4.png",
		},
	}

	validateInput(t, input, expected)
}

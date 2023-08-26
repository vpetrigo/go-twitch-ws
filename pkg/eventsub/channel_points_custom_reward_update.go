package eventsub

type ChannelPointsCustomRewardUpdateEvent struct {
	ID                                string              `json:"id"`                                    // The reward identifier.
	BroadcasterUserID                 string              `json:"broadcaster_user_id"`                   // The requested broadcaster ID.
	BroadcasterUserLogin              string              `json:"broadcaster_user_login"`                // The requested broadcaster login.
	BroadcasterUserName               string              `json:"broadcaster_user_name"`                 // The requested broadcaster display name.
	IsEnabled                         bool                `json:"is_enabled"`                            // Is the reward currently enabled.
	IsPaused                          bool                `json:"is_paused"`                             // Is the reward currently paused.
	IsInStock                         bool                `json:"is_in_stock"`                           // Is the reward currently in stock.
	Title                             string              `json:"title"`                                 // The reward title.
	Cost                              int                 `json:"cost"`                                  // The reward cost.
	Prompt                            string              `json:"prompt"`                                // The reward description.
	IsUserInputRequired               bool                `json:"is_user_input_required"`                // Does the viewer need to enter information when redeeming the reward.
	ShouldRedemptionsSkipRequestQueue bool                `json:"should_redemptions_skip_request_queue"` // Should redemptions be set tofulfilledstatus immediately when redeemed and skip the request queue instead of the normalunfulfilledstatus.
	MaxPerStream                      MaxPerStream        `json:"max_per_stream"`                        // Whether a maximum per stream is enabled and what the maximum is.
	MaxPerUserPerStream               MaxPerUserPerStream `json:"max_per_user_per_stream"`               // Whether a maximum per user per stream is enabled and what the maximum is.
	BackgroundColor                   string              `json:"background_color"`                      // Custom background color for the reward.
	Image                             Image               `json:"image"`                                 // Set of custom images of 1x, 2x and 4x sizes for the reward.
	DefaultImage                      interface{}         `json:"default_image"`                         // Set of default images of 1x, 2x and 4x sizes for the reward.
	GlobalCooldown                    GlobalCooldown      `json:"global_cooldown"`                       // Whether a cooldown is enabled and what the cooldown is in seconds.
	CooldownExpiresAt                 string              `json:"cooldown_expires_at"`                   // Timestamp of the cooldown expiration.
	RedemptionsRedeemedCurrentStream  int                 `json:"redemptions_redeemed_current_stream"`   // The number of redemptions redeemed during the current live stream.
}

type ChannelPointsCustomRewardUpdateEventCondition struct{}

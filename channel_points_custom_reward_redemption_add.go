package twitchws

type ChannelPointsCustomRewardRedemptionAddEvent struct {
	Id                   string      `json:"id"`                     // The redemption identifier.
	BroadcasterUserId    string      `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string      `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string      `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	UserId               string      `json:"user_id"`                // User ID of the user that redeemed the reward.
	UserLogin            string      `json:"user_login"`             // Login of the user that redeemed the reward.
	UserName             string      `json:"user_name"`              // Display name of the user that redeemed the reward.
	UserInput            string      `json:"user_input"`             // The user input provided.
	Status               string      `json:"status"`                 // Defaults tounfulfilled.
	Reward               interface{} `json:"reward"`                 // Basic information about the reward that was redeemed, at the time it was redeemed.
	RedeemedAt           string      `json:"redeemed_at"`            // RFC3339 timestamp of when the reward was redeemed.
}

type ChannelPointsCustomRewardRedemptionAddEventCondition struct{}

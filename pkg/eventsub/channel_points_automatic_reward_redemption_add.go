package eventsub

type ChannelPointsAutomaticRewardRedemptionAddEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The ID of the channel where the reward was redeemed.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The login of the channel where the reward was redeemed.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The display name of the channel where the reward was redeemed.
	UserID               string `json:"user_id"`                // The ID of the redeeming user.
	UserLogin            string `json:"user_login"`             // The login of the redeeming user.
	UserName             string `json:"user_name"`              // The display name of the redeeming user.
	ID                   string `json:"id"`                     // The ID of the Redemption.
	Reward               Reward `json:"reward"`                 // An object that contains the reward information.
	Message              string `json:"message"`                // An object that contains the user message and emote information needed to recreate the message.
	UserInput            string `json:"user_input"`             // Optional.
	RedeemedAt           string `json:"redeemed_at"`            // The UTC date and time (in RFC3339 format) of when the reward was redeemed.
}

package twitchws

type ChannelSubscriptionGiftEvent struct {
	UserID               string `json:"user_id"`                // The user ID of the user who sent the subscription gift.
	UserLogin            string `json:"user_login"`             // The user login of the user who sent the gift.
	UserName             string `json:"user_name"`              // The user display name of the user who sent the gift.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster display name.
	Total                int    `json:"total"`                  // The number of subscriptions in the subscription gift.
	Tier                 string `json:"tier"`                   // The tier of subscriptions in the subscription gift.
	CumulativeTotal      int    `json:"cumulative_total"`       // The number of subscriptions gifted by this user in the channel.
	IsAnonymous          bool   `json:"is_anonymous"`           // Whether the subscription gift was anonymous.
}

type ChannelSubscriptionGiftEventCondition struct{}

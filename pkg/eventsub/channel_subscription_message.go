package eventsub

type ChannelSubscriptionMessageEvent struct {
	UserID               string  `json:"user_id"`                // The user ID of the user who sent a resubscription chat message.
	UserLogin            string  `json:"user_login"`             // The user login of the user who sent a resubscription chat message.
	UserName             string  `json:"user_name"`              // The user display name of the user who a resubscription chat message.
	BroadcasterUserID    string  `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserLogin string  `json:"broadcaster_user_login"` // The broadcaster login.
	BroadcasterUserName  string  `json:"broadcaster_user_name"`  // The broadcaster display name.
	Tier                 string  `json:"tier"`                   // The tier of the user’s subscription.
	Message              Message `json:"message"`                // An object that contains the resubscription message and emote information needed to recreate the message.
	CumulativeMonths     int     `json:"cumulative_months"`      // The total number of months the user has been subscribed to the channel.
	StreakMonths         int     `json:"streak_months"`          // The number of consecutive months the user’s current subscription has been active.
	DurationMonths       int     `json:"duration_months"`        // The month duration of the subscription.
}

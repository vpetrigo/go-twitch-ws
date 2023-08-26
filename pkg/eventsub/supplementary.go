package eventsub

type Amount struct {
	Value         int    `json:"value"`          // The monetary amount.
	DecimalPlaces int    `json:"decimal_places"` // The number of decimal places used by the currency.
	Currency      string `json:"currency"`       // The ISO-4217 three-letter currency code that identifies the type of currency in value.
}

type BitsVoting struct {
	IsEnabled     bool `json:"is_enabled"`      // Not used; will be set to false.
	AmountPerVote int  `json:"amount_per_vote"` // Not used; will be set to 0.
}

type ChannelPointsVoting struct {
	IsEnabled     bool `json:"is_enabled"`      // Indicates if Channel Points can be used for voting.
	AmountPerVote int  `json:"amount_per_vote"` // Number of Channel Points required to vote once with Channel Points.
}

type Choices struct {
	ID                 string `json:"id"`                   // ID for the choice.
	Title              string `json:"title"`                // Text displayed for the choice.
	BitsVotes          int    `json:"bits_votes"`           // Not used; will be set to 0.
	ChannelPointsVotes int    `json:"channel_points_votes"` // Number of votes received via Channel Points.
	Votes              int    `json:"votes"`                // Total number of votes received for the choice across all methods of voting.
}

type CurrentAmount struct {
	Value         int    `json:"value"`          // The monetary amount.
	DecimalPlaces int    `json:"decimal_places"` // The number of decimal places used by the currency.
	Currency      string `json:"currency"`       // The ISO-4217 three-letter currency code that identifies the type of currency in value.
}

type Emotes struct {
	Begin int    `json:"begin"` // The index of where the Emote starts in the text.
	End   int    `json:"end"`   // The index of where the Emote ends in the text.
	ID    string `json:"id"`    // The emote ID.
}

type GlobalCooldown struct {
	IsEnabled bool `json:"is_enabled"` // Is the setting enabled.
	Seconds   int  `json:"seconds"`    // The cooldown in seconds.
}

type Image struct {
	URL1X string `json:"url_1x"` // URL for the image at 1x size.
	URL2X string `json:"url_2x"` // URL for the image at 2x size.
	URL4X string `json:"url_4x"` // URL for the image at 4x size.
}

type LastContribution struct {
	UserID    string `json:"user_id"`    // The ID of the user that made the contribution.
	UserLogin string `json:"user_login"` // The user’s login name.
	UserName  string `json:"user_name"`  // The user’s display name.
	Type      string `json:"type"`       // The contribution method used.
	Total     int    `json:"total"`      // The total amount contributed.
}

type MaxPerStream struct {
	IsEnabled bool `json:"is_enabled"` // Is the setting enabled.
	Value     int  `json:"value"`      // The max per stream limit.
}

type MaxPerUserPerStream struct {
	IsEnabled bool `json:"is_enabled"` // Is the setting enabled.
	Value     int  `json:"value"`      // The max per user per stream limit.
}

type Message struct {
	Text   string `json:"text"`   // The text of the resubscription chat message.
	Emotes Emotes `json:"emotes"` // An array that includes the emote ID and start and end positions for where the emote appears in the text.
}

type Outcomes struct {
	ID            string        `json:"id"`             // The outcome ID.
	Title         string        `json:"title"`          // The outcome title.
	Color         string        `json:"color"`          // The color for the outcome.
	Users         int           `json:"users"`          // The number of users who used Channel Points on this outcome.
	ChannelPoints int           `json:"channel_points"` // The total number of Channel Points used on this outcome.
	TopPredictors TopPredictors `json:"top_predictors"` // An array of users who used the most Channel Points on this outcome.
}

type Product struct {
	Name          string `json:"name"`           // Product name.
	Bits          int    `json:"bits"`           // Bits involved in the transaction.
	Sku           string `json:"sku"`            // Unique identifier for the product acquired.
	InDevelopment bool   `json:"in_development"` // Flag indicating if the product is in development.
}

type Reward struct {
	ID     string `json:"id"`     // The reward identifier.
	Title  string `json:"title"`  // The reward name.
	Cost   int    `json:"cost"`   // The reward cost.
	Prompt string `json:"prompt"` // The reward description.
}

type TargetAmount struct {
	Value         int    `json:"value"`          // The monetary amount.
	DecimalPlaces int    `json:"decimal_places"` // The number of decimal places used by the currency.
	Currency      string `json:"currency"`       // The ISO-4217 three-letter currency code that identifies the type of currency in value.
}

type TopContributions struct {
	UserID    string `json:"user_id"`    // The ID of the user that made the contribution.
	UserLogin string `json:"user_login"` // The user’s login name.
	UserName  string `json:"user_name"`  // The user’s display name.
	Type      string `json:"type"`       // The contribution method used.
	Total     int    `json:"total"`      // The total amount contributed.
}

type TopPredictors struct {
	UserID            string `json:"user_id"`             // The ID of the user.
	UserLogin         string `json:"user_login"`          // The login of the user.
	UserName          string `json:"user_name"`           // The display name of the user.
	ChannelPointsWon  int    `json:"channel_points_won"`  // The number of Channel Points won.
	ChannelPointsUsed int    `json:"channel_points_used"` // The number of Channel Points used to participate in the Prediction.
}

type dropEntitlementGrantEventData struct {
	OrganizationID string `json:"organization_id"` // The ID of the organization that owns the game that has Drops enabled.
	CategoryID     string `json:"category_id"`     // Twitch category ID of the game that was being played when this benefit was entitled.
	CategoryName   string `json:"category_name"`   // The category name.
	CampaignID     string `json:"campaign_id"`     // The campaign this entitlement is associated with.
	UserID         string `json:"user_id"`         // Twitch user ID of the user who was granted the entitlement.
	UserName       string `json:"user_name"`       // The user display name of the user who was granted the entitlement.
	UserLogin      string `json:"user_login"`      // The user login of the user who was granted the entitlement.
	EntitlementID  string `json:"entitlement_id"`  // Unique identifier of the entitlement.
	BenefitID      string `json:"benefit_id"`      // Identifier of the Benefit.
	CreatedAt      string `json:"created_at"`      // UTC timestamp in ISO format when this entitlement was granted on Twitch.
}

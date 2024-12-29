package eventsub

type Amount struct {
	Value         int    `json:"value"`          // The monetary amount.
	DecimalPlaces int    `json:"decimal_places"` // The number of decimal places used by the currency.
	Currency      string `json:"currency"`       // The ISO-4217 three-letter currency code that identifies the type of currency in value.
}

type Announcement struct {
	Color string `json:"color"` // Color of the announcement.
}

type AutomodTerms struct {
	Action      string   `json:"action"`       // Either “add” or “remove”.
	List        string   `json:"list"`         // Either “blocked” or “permitted”.
	Terms       []string `json:"terms"`        // Terms being added or removed.
	FromAutomod bool     `json:"from_automod"` // Whether the terms were added due to an Automod message approve/deny action.
}

type Ban struct {
	UserID    string `json:"user_id"`    // The ID of the user being banned.
	UserLogin string `json:"user_login"` // The login of the user being banned.
	UserName  string `json:"user_name"`  // The user name of the user being banned.
	Reason    string `json:"reason"`     // Optional.
}

type BitsBadgeTier struct {
	Tier int `json:"tier"` // The tier of the Bits badge the user just earned.
}

type BitsVoting struct {
	IsEnabled     bool `json:"is_enabled"`      // Not used; will be set to false.
	AmountPerVote int  `json:"amount_per_vote"` // Not used; will be set to 0.
}

type ChannelPointsVoting struct {
	IsEnabled     bool `json:"is_enabled"`      // Indicates if Channel Points can be used for voting.
	AmountPerVote int  `json:"amount_per_vote"` // Number of Channel Points required to vote once with Channel Points.
}

type Cheer struct {
	Bits int `json:"bits"` // The amount of Bits the user cheered.
}

type Choices struct {
	ID                 string `json:"id"`                   // ID for the choice.
	Title              string `json:"title"`                // Text displayed for the choice.
	BitsVotes          int    `json:"bits_votes"`           // Not used; will be set to 0.
	ChannelPointsVotes int    `json:"channel_points_votes"` // Number of votes received via Channel Points.
	Votes              int    `json:"votes"`                // Total number of votes received for the choice across all methods of voting.
}

type CommunitySubGift struct {
	ID              string `json:"id"`               // The ID of the associated community gift.
	Total           int    `json:"total"`            // Number of subscriptions being gifted.
	SubTier         string `json:"sub_tier"`         // The type of subscription plan being used.
	CumulativeTotal int    `json:"cumulative_total"` // Optional.
}

type CurrentAmount struct {
	Value         int    `json:"value"`          // The monetary amount.
	DecimalPlaces int    `json:"decimal_places"` // The number of decimal places used by the currency.
	Currency      string `json:"currency"`       // The ISO-4217 three-letter currency code that identifies the type of currency in value.
}

type Delete struct {
	UserID      string `json:"user_id"`      // The ID of the user whose message is being deleted.
	UserLogin   string `json:"user_login"`   // The login of the user.
	UserName    string `json:"user_name"`    // The user name of the user.
	MessageID   string `json:"message_id"`   // The ID of the message being deleted.
	MessageBody string `json:"message_body"` // The message body of the message being deleted.
}

type Emotes struct {
	Begin int    `json:"begin"` // The index of where the Emote starts in the text.
	End   int    `json:"end"`   // The index of where the Emote ends in the text.
	ID    string `json:"id"`    // The emote ID.
}

type Followers struct {
	FollowDurationMinutes int `json:"follow_duration_minutes"` // The length of time, in minutes, that the followers must have followed the broadcaster to participate in the chat room.
}

type GiftPaidUpgrade struct {
	GifterIsAnonymous bool   `json:"gifter_is_anonymous"` // Whether the gift was given anonymously.
	GifterUserID      string `json:"gifter_user_id"`      // Optional.
	GifterUserName    string `json:"gifter_user_name"`    // Optional.
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
	Text   string   `json:"text"`   // The text of the resubscription chat message.
	Emotes []Emotes `json:"emotes"` // An array that includes the emote ID and start and end positions for where the emote appears in the text.
}

type Mod struct {
	UserID    string `json:"user_id"`    // The ID of the user gaining mod status.
	UserLogin string `json:"user_login"` // The login of the user gaining mod status.
	UserName  string `json:"user_name"`  // The user name of the user gaining mod status.
}

type Outcomes struct {
	ID            string          `json:"id"`             // The outcome ID.
	Title         string          `json:"title"`          // The outcome title.
	Color         string          `json:"color"`          // The color for the outcome.
	Users         int             `json:"users"`          // The number of users who used Channel Points on this outcome.
	ChannelPoints int             `json:"channel_points"` // The total number of Channel Points used on this outcome.
	TopPredictors []TopPredictors `json:"top_predictors"` // An array of users who used the most Channel Points on this outcome.
}

type Participants struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The User ID of the participant channel.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The display name of the participant channel.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The user login of the participant channel.
}

type PayItForward struct {
	GifterIsAnonymous bool   `json:"gifter_is_anonymous"` // Whether the gift was given anonymously.
	GifterUserID      string `json:"gifter_user_id"`      // The user ID of the user who gifted the subscription.
	GifterUserName    string `json:"gifter_user_name"`    // Optional.
	GifterUserLogin   string `json:"gifter_user_login"`   // The user login of the user who gifted the subscription.
}

type PrimePaidUpgrade struct {
	SubTier string `json:"sub_tier"` // The type of subscription plan being used.
}

type Product struct {
	Name          string `json:"name"`           // Product name.
	Bits          int    `json:"bits"`           // Bits involved in the transaction.
	Sku           string `json:"sku"`            // Unique identifier for the product acquired.
	InDevelopment bool   `json:"in_development"` // Flag indicating if the product is in development.
}

type Raid struct {
	UserID          string `json:"user_id"`           // The user ID of the broadcaster raiding this channel.
	UserName        string `json:"user_name"`         // The user name of the broadcaster raiding this channel.
	UserLogin       string `json:"user_login"`        // The login name of the broadcaster raiding this channel.
	ViewerCount     int    `json:"viewer_count"`      // The number of viewers raiding this channel from the broadcaster’s channel.
	ProfileImageURL string `json:"profile_image_url"` // Profile image URL of the broadcaster raiding this channel.
}

type Reply struct {
	ParentMessageID   string `json:"parent_message_id"`   // An ID that uniquely identifies the parent message that this message is replying to.
	ParentMessageBody string `json:"parent_message_body"` // The message body of the parent message.
	ParentUserID      string `json:"parent_user_id"`      // User ID of the sender of the parent message.
	ParentUserName    string `json:"parent_user_name"`    // User name of the sender of the parent message.
	ParentUserLogin   string `json:"parent_user_login"`   // User login of the sender of the parent message.
	ThreadMessageID   string `json:"thread_message_id"`   // An ID that identifies the parent message of the reply thread.
	ThreadUserID      string `json:"thread_user_id"`      // User ID of the sender of the thread’s parent message.
	ThreadUserName    string `json:"thread_user_name"`    // User name of the sender of the thread’s parent message.
	ThreadUserLogin   string `json:"thread_user_login"`   // User login of the sender of the thread’s parent message.
}

type Resub struct {
	CumulativeMonths  int    `json:"cumulative_months"`   // The total number of months the user has subscribed.
	DurationMonths    int    `json:"duration_months"`     // The number of months the subscription is for.
	StreakMonths      int    `json:"streak_months"`       // The total number of months the user has subscribed.
	SubTier           string `json:"sub_tier"`            // The type of subscription plan being used.
	IsPrime           bool   `json:"is_prime"`            // Optional.
	IsGift            bool   `json:"is_gift"`             // Whether or not the resub was a result of a gift.
	GifterIsAnonymous bool   `json:"gifter_is_anonymous"` // Optional.
	GifterUserID      string `json:"gifter_user_id"`      // The user ID of the subscription gifter.
	GifterUserName    string `json:"gifter_user_name"`    // The user name of the subscription gifter.
	GifterUserLogin   string `json:"gifter_user_login"`   // Optional.
}

type Reward struct {
	ID     string `json:"id"`     // The reward identifier.
	Title  string `json:"title"`  // The reward name.
	Cost   int    `json:"cost"`   // The reward cost.
	Prompt string `json:"prompt"` // The reward description.
}

type Slow struct {
	WaitTimeSeconds int `json:"wait_time_seconds"` // The amount of time, in seconds, that users need to wait between sending messages.
}

type SourceBadges struct {
	SetID string `json:"set_id"` // The ID that identifies this set of chat badges.
	ID    string `json:"id"`     // The ID that identifies this version of the badge.
	Info  string `json:"info"`   // Contains metadata related to the chat badges in the badges tag.
}

type Sub struct {
	SubTier        string `json:"sub_tier"`        // The type of subscription plan being used.
	IsPrime        bool   `json:"is_prime"`        // Indicates if the subscription was obtained through Amazon Prime.
	DurationMonths int    `json:"duration_months"` // The number of months the subscription is for.
}

type SubGift struct {
	DurationMonths     int    `json:"duration_months"`      // The number of months the subscription is for.
	CumulativeTotal    int    `json:"cumulative_total"`     // Optional.
	RecipientUserID    string `json:"recipient_user_id"`    // The user ID of the subscription gift recipient.
	RecipientUserName  string `json:"recipient_user_name"`  // The user name of the subscription gift recipient.
	RecipientUserLogin string `json:"recipient_user_login"` // The user login of the subscription gift recipient.
	SubTier            string `json:"sub_tier"`             // The type of subscription plan being used.
	CommunityGiftID    string `json:"community_gift_id"`    // Optional.
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

type Timeout struct {
	UserID    string `json:"user_id"`    // The ID of the user being timed out.
	UserLogin string `json:"user_login"` // The login of the user being timed out.
	UserName  string `json:"user_name"`  // The user name of the user being timed out.
	Reason    string `json:"reason"`     // Optional.
	ExpiresAt string `json:"expires_at"` // The time at which the timeout ends.
}

type TopPredictors struct {
	UserID            string `json:"user_id"`             // The ID of the user.
	UserLogin         string `json:"user_login"`          // The login of the user.
	UserName          string `json:"user_name"`           // The display name of the user.
	ChannelPointsWon  int    `json:"channel_points_won"`  // The number of Channel Points won.
	ChannelPointsUsed int    `json:"channel_points_used"` // The number of Channel Points used to participate in the Prediction.
}

type Unban struct {
	UserID    string `json:"user_id"`    // The ID of the user being unbanned.
	UserLogin string `json:"user_login"` // The login of the user being unbanned.
	UserName  string `json:"user_name"`  // The user name of the user being unbanned.
}

type UnbanRequest struct {
	IsApproved       bool   `json:"is_approved"`       // Whether or not the unban request was approved or denied.
	UserID           string `json:"user_id"`           // The ID of the banned user.
	UserLogin        string `json:"user_login"`        // The login of the user.
	UserName         string `json:"user_name"`         // The user name of the user.
	ModeratorMessage string `json:"moderator_message"` // The message included by the moderator explaining their approval or denial.
}

type Unmod struct {
	UserID    string `json:"user_id"`    // The ID of the user losing mod status.
	UserLogin string `json:"user_login"` // The login of the user losing mod status.
	UserName  string `json:"user_name"`  // The user name of the user losing mod status.
}

type Unraid struct {
	UserID    string `json:"user_id"`    // The ID of the user no longer being raided.
	UserLogin string `json:"user_login"` // The login of the user no longer being raided.
	UserName  string `json:"user_name"`  // The user name of the no longer user raided.
}

type Untimeout struct {
	UserID    string `json:"user_id"`    // The ID of the user being untimed out.
	UserLogin string `json:"user_login"` // The login of the user being untimed out.
	UserName  string `json:"user_name"`  // The user name of the user untimed out.
}

type Unvip struct {
	UserID    string `json:"user_id"`    // The ID of the user losing VIP status.
	UserLogin string `json:"user_login"` // The login of the user losing VIP status.
	UserName  string `json:"user_name"`  // The user name of the user losing VIP status.
}

type Vip struct {
	UserID    string `json:"user_id"`    // The ID of the user gaining VIP status.
	UserLogin string `json:"user_login"` // The login of the user gaining VIP status.
	UserName  string `json:"user_name"`  // The user name of the user gaining VIP status.
}

type Whisper struct {
	Text string `json:"text"` // The body of the whisper message.
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

package twitchws

type ChannelBanEvent struct {
	UserID               string `json:"user_id"`                // The user ID for the user who was banned on the specified channel.
	UserLogin            string `json:"user_login"`             // The user login for the user who was banned on the specified channel.
	UserName             string `json:"user_name"`              // The user display name for the user who was banned on the specified channel.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	ModeratorUserID      string `json:"moderator_user_id"`      // The user ID of the issuer of the ban.
	ModeratorUserLogin   string `json:"moderator_user_login"`   // The user login of the issuer of the ban.
	ModeratorUserName    string `json:"moderator_user_name"`    // The user name of the issuer of the ban.
	Reason               string `json:"reason"`                 // The reason behind the ban.
	BannedAt             string `json:"banned_at"`              // The UTC date and time (in RFC3339 format) of when the user was banned or put in a timeout.
	EndsAt               string `json:"ends_at"`                // The UTC date and time (in RFC3339 format) of when the timeout ends.
	IsPermanent          bool   `json:"is_permanent"`           // Indicates whether the ban is permanent (true) or a timeout (false).
}

type ChannelBanEventCondition struct{}

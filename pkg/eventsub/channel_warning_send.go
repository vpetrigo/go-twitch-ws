package eventsub

type ChannelWarningSendEvent struct {
	BroadcasterUserID    string      `json:"broadcaster_user_id"`    // The user ID of the broadcaster.
	BroadcasterUserLogin string      `json:"broadcaster_user_login"` // The login of the broadcaster.
	BroadcasterUserName  string      `json:"broadcaster_user_name"`  // The user name of the broadcaster.
	ModeratorUserID      string      `json:"moderator_user_id"`      // The user ID of the moderator who sent the warning.
	ModeratorUserLogin   string      `json:"moderator_user_login"`   // The login of the moderator.
	ModeratorUserName    string      `json:"moderator_user_name"`    // The user name of the moderator.
	UserID               string      `json:"user_id"`                // The ID of the user being warned.
	UserLogin            string      `json:"user_login"`             // The login of the user being warned.
	UserName             string      `json:"user_name"`              // The user name of the user being.
	Reason               string      `json:"reason"`                 // Optional.
	ChatRulesCited       interface{} `json:"chat_rules_cited"`       // Optional.
}

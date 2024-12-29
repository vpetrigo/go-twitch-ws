package eventsub

type AutomodTermsUpdateEvent struct {
	BroadcasterUserID    string   `json:"broadcaster_user_id"`    // The ID of the broadcaster specified in the request.
	BroadcasterUserLogin string   `json:"broadcaster_user_login"` // The login of the broadcaster specified in the request.
	BroadcasterUserName  string   `json:"broadcaster_user_name"`  // The user name of the broadcaster specified in the request.
	ModeratorUserID      string   `json:"moderator_user_id"`      // The ID of the moderator who changed the channel settings.
	ModeratorUserLogin   string   `json:"moderator_user_login"`   // The moderator’s login.
	ModeratorUserName    string   `json:"moderator_user_name"`    // The moderator’s user name.
	Action               string   `json:"action"`                 // The status change applied to the terms.
	FromAutomod          bool     `json:"from_automod"`           // Indicates whether this term was added due to an Automod message approve/deny action.
	Terms                []string `json:"terms"`                  // The list of terms that had a status change.
}

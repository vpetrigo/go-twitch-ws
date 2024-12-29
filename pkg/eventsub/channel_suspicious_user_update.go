package eventsub

type ChannelSuspiciousUserUpdateEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The ID of the channel where the treatment for a suspicious user was updated.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The display name of the channel where the treatment for a suspicious user was updated.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The Login of the channel where the treatment for a suspicious user was updated.
	ModeratorUserID      string `json:"moderator_user_id"`      // The ID of the moderator that updated the treatment for a suspicious user.
	ModeratorUserName    string `json:"moderator_user_name"`    // The display name of the moderator that updated the treatment for a suspicious user.
	ModeratorUserLogin   string `json:"moderator_user_login"`   // The login of the moderator that updated the treatment for a suspicious user.
	UserID               string `json:"user_id"`                // The ID of the suspicious user whose treatment was updated.
	UserName             string `json:"user_name"`              // The display name of the suspicious user whose treatment was updated.
	UserLogin            string `json:"user_login"`             // The login of the suspicious user whose treatment was updated.
	LowTrustStatus       string `json:"low_trust_status"`       // The status set for the suspicious user.
}

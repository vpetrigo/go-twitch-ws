package twitchws

type ChannelUnbanEvent struct {
	UserId               string `json:"user_id"`                // The user id for the user who was unbanned on the specified channel.
	UserLogin            string `json:"user_login"`             // The user login for the user who was unbanned on the specified channel.
	UserName             string `json:"user_name"`              // The user display name for the user who was unbanned on the specified channel.
	BroadcasterUserId    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	ModeratorUserId      string `json:"moderator_user_id"`      // The user ID of the issuer of the unban.
	ModeratorUserLogin   string `json:"moderator_user_login"`   // The user login of the issuer of the unban.
	ModeratorUserName    string `json:"moderator_user_name"`    // The user name of the issuer of the unban.
}

type ChannelUnbanEventCondition struct{}

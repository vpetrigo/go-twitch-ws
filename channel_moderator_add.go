package twitchws

type ChannelModeratorAddEvent struct {
	BroadcasterUserId    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	UserId               string `json:"user_id"`                // The user ID of the new moderator.
	UserLogin            string `json:"user_login"`             // The user login of the new moderator.
	UserName             string `json:"user_name"`              // The display name of the new moderator.
}

type ChannelModeratorAddEventCondition struct{}

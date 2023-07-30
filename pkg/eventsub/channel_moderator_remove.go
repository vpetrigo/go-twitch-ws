package eventsub

type ChannelModeratorRemoveEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	UserID               string `json:"user_id"`                // The user ID of the removed moderator.
	UserLogin            string `json:"user_login"`             // The user login of the removed moderator.
	UserName             string `json:"user_name"`              // The display name of the removed moderator.
}

type ChannelModeratorRemoveEventCondition struct{}

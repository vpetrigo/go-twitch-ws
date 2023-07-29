package twitchws

// ChannelFollowEvent Contains the user ID and username of the follower
// and the broadcaster user ID and broadcaster username.
//
// The `channel.follow` subscription type sends a notification when a specified channel receives a follow.
// **Authorization:** Must have moderator:read:followers scope.
type ChannelFollowEvent struct {
	UserID               string `json:"user_id"`                // The user ID for the user now following the specified channel.
	UserLogin            string `json:"user_login"`             // The user login for the user now following the specified channel.
	UserName             string `json:"user_name"`              // The user display name for the user now following the specified channel.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	FollowedAt           string `json:"followed_at"`            // RFC3339 timestamp of when the follow occurred.
}

type ChannelFollowEventCondition struct {
	BroadcasterUserID string `json:"broadcaster_user_id"` // The broadcaster user ID for the channel you want to get follow notifications for.
	ModeratorUserID   string `json:"moderator_user_id"`   // The ID of the moderator of the channel you want to get follow notifications for. If you have authorization from the broadcaster rather than a moderator, specify the broadcaster’s user ID here.
}

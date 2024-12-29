package eventsub

type ChannelUnbanRequestCreateEvent struct {
	ID                   string `json:"id"`                     // The ID of the unban request.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster’s user ID for the channel the unban request was created for.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s login name.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s display name.
	UserID               string `json:"user_id"`                // User ID of user that is requesting to be unbanned.
	UserLogin            string `json:"user_login"`             // The user’s login name.
	UserName             string `json:"user_name"`              // The user’s display name.
	Text                 string `json:"text"`                   // Message sent in the unban request.
	CreatedAt            string `json:"created_at"`             // The UTC timestamp (in RFC3339 format) of when the unban request was created.
}

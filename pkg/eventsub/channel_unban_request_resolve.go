package eventsub

type ChannelUnbanRequestResolveEvent struct {
	ID                   string `json:"id"`                     // The ID of the unban request.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster’s user ID for the channel the unban request was updated for.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s login name.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s display name.
	ModeratorID          string `json:"moderator_id"`           // Optional.
	ModeratorLogin       string `json:"moderator_login"`        // Optional.
	ModeratorName        string `json:"moderator_name"`         // Optional.
	UserID               string `json:"user_id"`                // User ID of user that requested to be unbanned.
	UserLogin            string `json:"user_login"`             // The user’s login name.
	UserName             string `json:"user_name"`              // The user’s display name.
	ResolutionText       string `json:"resolution_text"`        // Optional.
	Status               string `json:"status"`                 // Dictates whether the unban request was approved or denied.
}

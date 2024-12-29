package eventsub

type ChannelAdBreakBeginEvent struct {
	DurationSeconds      int    `json:"duration_seconds"`       // Length in seconds of the mid-roll ad break requested.
	StartedAt            string `json:"started_at"`             // The UTC timestamp of when the ad break began, in RFC3339 format.
	IsAutomatic          bool   `json:"is_automatic"`           // Indicates if the ad was automatically scheduled via Ads Manager.
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster’s user ID for the channel the ad was run on.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster’s user login for the channel the ad was run on.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster’s user display name for the channel the ad was run on.
	RequesterUserID      string `json:"requester_user_id"`      // The ID of the user that requested the ad.
	RequesterUserLogin   string `json:"requester_user_login"`   // The login of the user that requested the ad.
	RequesterUserName    string `json:"requester_user_name"`    // The display name of the user that requested the ad.
}

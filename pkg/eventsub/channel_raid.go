package eventsub

type ChannelRaidEvent struct {
	FromBroadcasterUserID    string `json:"from_broadcaster_user_id"`    // The broadcaster ID that created the raid.
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"` // The broadcaster login that created the raid.
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`  // The broadcaster display name that created the raid.
	ToBroadcasterUserID      string `json:"to_broadcaster_user_id"`      // The broadcaster ID that received the raid.
	ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`   // The broadcaster login that received the raid.
	ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`    // The broadcaster display name that received the raid.
	Viewers                  int    `json:"viewers"`                     // The number of viewers in the raid.
}

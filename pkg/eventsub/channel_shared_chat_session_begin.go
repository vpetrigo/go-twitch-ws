package eventsub

type ChannelSharedChatSessionBeginEvent struct {
	SessionID                string       `json:"session_id"`                  // The unique identifier for the shared chat session.
	BroadcasterUserID        string       `json:"broadcaster_user_id"`         // The User ID of the channel in the subscription condition which is now active in the shared chat session.
	BroadcasterUserName      string       `json:"broadcaster_user_name"`       // The display name of the channel in the subscription condition which is now active in the shared chat session.
	BroadcasterUserLogin     string       `json:"broadcaster_user_login"`      // The user login of the channel in the subscription condition which is now active in the shared chat session.
	HostBroadcasterUserID    string       `json:"host_broadcaster_user_id"`    // The User ID of the host channel.
	HostBroadcasterUserName  string       `json:"host_broadcaster_user_name"`  // The display name of the host channel.
	HostBroadcasterUserLogin string       `json:"host_broadcaster_user_login"` // The user login of the host channel.
	Participants             Participants `json:"participants"`                // The list of participants in the session.
}

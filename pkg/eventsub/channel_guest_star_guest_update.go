package eventsub

type ChannelGuestStarGuestUpdateEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The non-host broadcaster user ID.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The non-host broadcaster display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The non-host broadcaster login.
	SessionID            string `json:"session_id"`             // ID representing the unique session that was started.
	ModeratorUserID      string `json:"moderator_user_id"`      // The user ID of the moderator who updated the guest’s state (could be the host).
	ModeratorUserName    string `json:"moderator_user_name"`    // The moderator display name.
	ModeratorUserLogin   string `json:"moderator_user_login"`   // The moderator login.
	GuestUserID          string `json:"guest_user_id"`          // The user ID of the guest who transitioned states in the session.
	GuestUserName        string `json:"guest_user_name"`        // The guest display name.
	GuestUserLogin       string `json:"guest_user_login"`       // The guest login.
	SlotID               string `json:"slot_id"`                // The ID of the slot assignment the guest is assigned to.
	State                string `json:"state"`                  // The current state of the user after the update has taken place.
	HostUserID           string `json:"host_user_id"`           // User ID of the host channel.
	HostUserName         string `json:"host_user_name"`         // The host display name.
	HostUserLogin        string `json:"host_user_login"`        // The host login.
	HostVideoEnabled     bool   `json:"host_video_enabled"`     // Flag that signals whether the host is allowing the slot’s video to be seen by participants within the session.
	HostAudioEnabled     bool   `json:"host_audio_enabled"`     // Flag that signals whether the host is allowing the slot’s audio to be heard by participants within the session.
	HostVolume           int    `json:"host_volume"`            // Value between 0-100 that represents the slot’s audio level as heard by participants within the session.
}

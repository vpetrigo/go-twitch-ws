package twitchws

type ChannelGuestStarSlotUpdateEvent struct {
	BroadcasterUserId    string      `json:"broadcaster_user_id"`    // User ID of the host channel.
	BroadcasterUserName  string      `json:"broadcaster_user_name"`  // The broadcaster display name.
	BroadcasterUserLogin string      `json:"broadcaster_user_login"` // The broadcaster login.
	SessionId            string      `json:"session_id"`             // ID representing the unique session the event took place in.
	ModeratorUserId      string      `json:"moderator_user_id"`      // The user ID of the moderator who modified the setting, if any.
	ModeratorUserName    string      `json:"moderator_user_name"`    // The moderator display name.
	ModeratorUserLogin   string      `json:"moderator_user_login"`   // The moderator login.
	GuestUserId          string      `json:"guest_user_id"`          // The user ID of the guest who is assigned to the slot.
	GuestUserName        string      `json:"guest_user_name"`        // The guest display name.
	GuestUserLogin       string      `json:"guest_user_login"`       // The guest login.
	SlotId               string      `json:"slot_id"`                // The ID of the slot where settings were updated.
	HostVideoEnabled     bool        `json:"host_video_enabled"`     // Flag that signals whether the host is allowing the slot’s video to be seen by participants within the session.
	HostAudioEnabled     bool        `json:"host_audio_enabled"`     // Flag that signals whether the host is allowing the slot’s audio to be heard by participants within the session.
	HostVolume           interface{} `json:"host_volume"`            // Value between 0-100 that represents the slot’s audio level as heard by participants within the session.
}

type ChannelGuestStarSlotUpdateEventCondition struct{}

package twitchws

type ChannelGuestStarSettingsUpdateEvent struct {
	BroadcasterUserId           string `json:"broadcaster_user_id"`             // User ID of the host channel.
	BroadcasterUserName         string `json:"broadcaster_user_name"`           // The broadcaster display name.
	BroadcasterUserLogin        string `json:"broadcaster_user_login"`          // The broadcaster login.
	IsModeratorSendLiveEnabled  bool   `json:"is_moderator_send_live_enabled"`  // Flag determining if Guest Star moderators have access to control whether a guest is live once assigned to a slot.
	SlotCount                   int    `json:"slot_count"`                      // Number of slots the Guest Star call interface will allow the host to add to a call.
	IsBrowserSourceAudioEnabled bool   `json:"is_browser_source_audio_enabled"` // Flag determining if browser sources subscribed to sessions on this channel should output audio.
	GroupLayout                 string `json:"group_layout"`                    // This setting determines how the guests within a session should be laid out within a group browser source.
}

type ChannelGuestStarSettingsUpdateEventCondition struct{}

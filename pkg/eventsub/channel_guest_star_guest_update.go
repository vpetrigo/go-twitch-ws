package eventsub

type ChannelGuestStarGuestUpdateEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`    // The broadcaster user ID.
	BroadcasterUserName  string `json:"broadcaster_user_name"`  // The broadcaster display name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"` // The broadcaster login.
	SessionID            string `json:"session_id"`             // ID representing the unique session that was started.
	ModeratorUserID      string `json:"moderator_user_id"`      // The user ID of the moderator who updated the guest’s state (could be the host).
	ModeratorUserName    string `json:"moderator_user_name"`    // The moderator display name.
	ModeratorUserLogin   string `json:"moderator_user_login"`   // The moderator login.
	GuestUserID          string `json:"guest_user_id"`          // The user ID of the guest who transitioned states in the session.
	GuestUserName        string `json:"guest_user_name"`        // The guest display name.
	GuestUserLogin       string `json:"guest_user_login"`       // The guest login.
	SlotID               string `json:"slot_id"`                // The ID of the slot assignment the guest is assigned to.
	State                string `json:"state"`                  // The current state of the user after the update has taken place.
}

type ChannelGuestStarGuestUpdateEventCondition struct{}

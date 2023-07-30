package eventsub

type UserUpdateEvent struct {
	UserID        string `json:"user_id"`        // The user’s user id.
	UserLogin     string `json:"user_login"`     // The user’s user login.
	UserName      string `json:"user_name"`      // The user’s user display name.
	Email         string `json:"email"`          // The user’s email address.
	EmailVerified bool   `json:"email_verified"` // A Boolean value that determines whether Twitch has verified the user’s email address.
	Description   string `json:"description"`    // The user’s description.
}

type UserUpdateEventCondition struct{}

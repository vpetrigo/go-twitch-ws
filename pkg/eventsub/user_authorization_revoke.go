package eventsub

type UserAuthorizationRevokeEvent struct {
	ClientID  string `json:"client_id"`  // The client_id of the application with revoked user access.
	UserID    string `json:"user_id"`    // The user id for the user who has revoked authorization for your client id.
	UserLogin string `json:"user_login"` // The user login for the user who has revoked authorization for your client id.
	UserName  string `json:"user_name"`  // The user display name for the user who has revoked authorization for your client id.
}

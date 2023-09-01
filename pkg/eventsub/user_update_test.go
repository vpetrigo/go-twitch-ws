package eventsub

import (
	"testing"
)

func TestUserUpdate(t *testing.T) {
	input := `{
        "user_id": "1337",
        "user_login": "cool_user",
        "user_name": "Cool_User",
        "email": "user@email.com", 
        "email_verified": true,
        "description": "cool description"
    }`
	expected := UserUpdateEvent{
		UserID:        "1337",
		UserLogin:     "cool_user",
		UserName:      "Cool_User",
		Email:         "user@email.com",
		EmailVerified: true,
		Description:   "cool description",
	}

	validateInput(t, input, expected)
}

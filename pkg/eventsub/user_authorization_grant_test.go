package eventsub

import (
	"testing"
)

func TestUserAuthorizationGrant(t *testing.T) {
	input := `{
        "client_id": "crq72vsaoijkc83xx42hz6i37",
        "user_id": "1337",
        "user_login": "cool_user",
        "user_name": "Cool_User"
    }`
	expected := UserAuthorizationGrantEvent{
		ClientID:  "crq72vsaoijkc83xx42hz6i37",
		UserID:    "1337",
		UserLogin: "cool_user",
		UserName:  "Cool_User",
	}
	validateInput(t, input, expected)
}

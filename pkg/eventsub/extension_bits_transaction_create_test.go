package eventsub

import "testing"

func TestExtensionBitsTransactionCreate(t *testing.T) {
	input := `{
        "id": "bits-tx-id",
        "extension_client_id": "deadbeef",
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "user_name": "Coolest_User",
        "user_login": "coolest_user",
        "user_id": "1236",
        "product": {
            "name": "great_product",
            "sku": "skuskusku",
            "bits": 1234,
            "in_development": false
        }
    }`
	expected := ExtensionBitsTransactionCreateEvent{
		ID:                   "bits-tx-id",
		ExtensionClientID:    "deadbeef",
		BroadcasterUserID:    "1337",
		BroadcasterUserLogin: "cool_user",
		BroadcasterUserName:  "Cool_User",
		UserName:             "Coolest_User",
		UserLogin:            "coolest_user",
		UserID:               "1236",
		Product: Product{
			Name:          "great_product",
			Sku:           "skuskusku",
			Bits:          1234,
			InDevelopment: false,
		},
	}

	validateInput(t, input, expected)
}

package twitchws

type ExtensionBitsTransactionCreateEvent struct {
	ExtensionClientId    string      `json:"extension_client_id"`    // Client ID of the extension.
	Id                   string      `json:"id"`                     // Transaction ID.
	BroadcasterUserId    string      `json:"broadcaster_user_id"`    // The transaction’s broadcaster ID.
	BroadcasterUserLogin string      `json:"broadcaster_user_login"` // The transaction’s broadcaster login.
	BroadcasterUserName  string      `json:"broadcaster_user_name"`  // The transaction’s broadcaster display name.
	UserId               string      `json:"user_id"`                // The transaction’s user ID.
	UserLogin            string      `json:"user_login"`             // The transaction’s user login.
	UserName             string      `json:"user_name"`              // The transaction’s user display name.
	Product              interface{} `json:"product"`                // Additional extension product information.
}

type ExtensionBitsTransactionCreateEventCondition struct{}

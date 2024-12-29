package eventsub

type WhisperReceivedEvent struct {
	FromUserID    string  `json:"from_user_id"`    // The ID of the user sending the message.
	FromUserName  string  `json:"from_user_name"`  // The name of the user sending the message.
	FromUserLogin string  `json:"from_user_login"` // The login of the user sending the message.
	ToUserID      string  `json:"to_user_id"`      // The ID of the user receiving the message.
	ToUserName    string  `json:"to_user_name"`    // The name of the user receiving the message.
	ToUserLogin   string  `json:"to_user_login"`   // The login of the user receiving the message.
	WhisperID     string  `json:"whisper_id"`      // The whisper ID.
	Whisper       Whisper `json:"whisper"`         // Object containing whisper information.
}

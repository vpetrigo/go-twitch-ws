package eventsub

type ChannelSuspiciousUserMessageEvent struct {
	BroadcasterUserID    string      `json:"broadcaster_user_id"`    // The ID of the channel where the treatment for a suspicious user was updated.
	BroadcasterUserName  string      `json:"broadcaster_user_name"`  // The display name of the channel where the treatment for a suspicious user was updated.
	BroadcasterUserLogin string      `json:"broadcaster_user_login"` // The login of the channel where the treatment for a suspicious user was updated.
	UserID               string      `json:"user_id"`                // The user ID of the user that sent the message.
	UserName             string      `json:"user_name"`              // The user name of the user that sent the message.
	UserLogin            string      `json:"user_login"`             // The user login of the user that sent the message.
	LowTrustStatus       string      `json:"low_trust_status"`       // The status set for the suspicious user.
	SharedBanChannelIDs  interface{} `json:"shared_ban_channel_ids"` // A list of channel IDs where the suspicious user is also banned.
	Types                interface{} `json:"types"`                  // User types (if any) that apply to the suspicious user, can be “manual”, “ban_evader_detector”, or “shared_channel_ban”.
	BanEvasionEvaluation string      `json:"ban_evasion_evaluation"` // A ban evasion likelihood value (if any) that as been applied to the user automatically by Twitch, can be “unknown”, “possible”, or “likely”.
	Message              string      `json:"message"`                // The structured chat message.
}

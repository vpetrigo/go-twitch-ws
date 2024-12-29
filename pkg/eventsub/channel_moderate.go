package eventsub

type ChannelModerateEvent struct {
	BroadcasterUserID          string       `json:"broadcaster_user_id"`           // The ID of the broadcaster.
	BroadcasterUserLogin       string       `json:"broadcaster_user_login"`        // The login of the broadcaster.
	BroadcasterUserName        string       `json:"broadcaster_user_name"`         // The user name of the broadcaster.
	SourceBroadcasterUserID    string       `json:"source_broadcaster_user_id"`    // The channel in which the action originally occurred.
	SourceBroadcasterUserLogin string       `json:"source_broadcaster_user_login"` // The channel in which the action originally occurred.
	SourceBroadcasterUserName  string       `json:"source_broadcaster_user_name"`  // The channel in which the action originally occurred.
	ModeratorUserID            string       `json:"moderator_user_id"`             // The ID of the moderator who performed the action.
	ModeratorUserLogin         string       `json:"moderator_user_login"`          // The login of the moderator.
	ModeratorUserName          string       `json:"moderator_user_name"`           // The user name of the moderator.
	Action                     string       `json:"action"`                        // The type of action: Possible values are: li.
	Followers                  Followers    `json:"followers"`                     // Optional.
	Slow                       Slow         `json:"slow"`                          // Optional.
	Vip                        Vip          `json:"vip"`                           // Optional.
	Unvip                      Unvip        `json:"unvip"`                         // Optional.
	Mod                        Mod          `json:"mod"`                           // Optional.
	Unmod                      Unmod        `json:"unmod"`                         // Optional.
	Ban                        Ban          `json:"ban"`                           // Optional.
	Unban                      Unban        `json:"unban"`                         // Optional.
	Timeout                    Timeout      `json:"timeout"`                       // Optional.
	Untimeout                  Untimeout    `json:"untimeout"`                     // Optional.
	Raid                       Raid         `json:"raid"`                          // Optional.
	Unraid                     Unraid       `json:"unraid"`                        // Optional.
	Delete                     Delete       `json:"delete"`                        // Optional.
	AutomodTerms               AutomodTerms `json:"automod_terms"`                 // Optional.
	UnbanRequest               UnbanRequest `json:"unban_request"`                 // Optional.
	SharedChatBan              interface{}  `json:"shared_chat_ban"`               // Optional.
	SharedChatUnban            interface{}  `json:"shared_chat_unban"`             // Optional.
	SharedChatTimeout          interface{}  `json:"shared_chat_timeout"`           // Optional.
	SharedChatUntimeout        interface{}  `json:"shared_chat_untimeout"`         // Optional.
	SharedChatDelete           interface{}  `json:"shared_chat_delete"`            // Optional.
}

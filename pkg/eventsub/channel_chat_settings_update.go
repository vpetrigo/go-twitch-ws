package eventsub

type ChannelChatSettingsUpdateEvent struct {
	BroadcasterUserID           string `json:"broadcaster_user_id"`            // The ID of the broadcaster specified in the request.
	BroadcasterUserLogin        string `json:"broadcaster_user_login"`         // The login of the broadcaster specified in the request.
	BroadcasterUserName         string `json:"broadcaster_user_name"`          // The user name of the broadcaster specified in the request.
	EmoteMode                   bool   `json:"emote_mode"`                     // A Boolean value that determines whether chat messages must contain only emotes.
	FollowerMode                bool   `json:"follower_mode"`                  // A Boolean value that determines whether the broadcaster restricts the chat room to followers only, based on how long they’ve followed.
	FollowerModeDurationMinutes int    `json:"follower_mode_duration_minutes"` // The length of time, in minutes, that the followers must have followed the broadcaster to participate in the chat room.
	SlowMode                    bool   `json:"slow_mode"`                      // A Boolean value that determines whether the broadcaster limits how often users in the chat room are allowed to send messages.
	SlowModeWaitTimeSeconds     int    `json:"slow_mode_wait_time_seconds"`    // The amount of time, in seconds, that users need to wait between sending messages.
	SubscriberMode              bool   `json:"subscriber_mode"`                // A Boolean value that determines whether only users that subscribe to the broadcaster’s channel can talk in the chat room.
	UniqueChatMode              bool   `json:"unique_chat_mode"`               // A Boolean value that determines whether the broadcaster requires users to post only unique messages in the chat room.
}

package twitchws

type ChannelPollBeginEvent struct {
	ID                   string      `json:"id"`                     // ID of the poll.
	BroadcasterUserID    string      `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string      `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string      `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Title                string      `json:"title"`                  // Question displayed for the poll.
	Choices              interface{} `json:"choices"`                // An array of choices for the poll.
	BitsVoting           interface{} `json:"bits_voting"`            // Not supported.
	ChannelPointsVoting  interface{} `json:"channel_points_voting"`  // The Channel Points voting settings for the poll.
	StartedAt            string      `json:"started_at"`             // The time the poll started.
	EndsAt               string      `json:"ends_at"`                // The time the poll will end.
}

type ChannelPollBeginEventCondition struct{}

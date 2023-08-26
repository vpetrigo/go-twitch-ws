package eventsub

type ChannelPollEndEvent struct {
	ID                   string              `json:"id"`                     // ID of the poll.
	BroadcasterUserID    string              `json:"broadcaster_user_id"`    // The requested broadcaster ID.
	BroadcasterUserLogin string              `json:"broadcaster_user_login"` // The requested broadcaster login.
	BroadcasterUserName  string              `json:"broadcaster_user_name"`  // The requested broadcaster display name.
	Title                string              `json:"title"`                  // Question displayed for the poll.
	Choices              Choices             `json:"choices"`                // An array of choices for the poll.
	BitsVoting           BitsVoting          `json:"bits_voting"`            // Not supported.
	ChannelPointsVoting  ChannelPointsVoting `json:"channel_points_voting"`  // The Channel Points voting settings for the poll.
	Status               string              `json:"status"`                 // The status of the poll.
	StartedAt            string              `json:"started_at"`             // The time the poll started.
	EndedAt              string              `json:"ended_at"`               // The time the poll ended.
}

type ChannelPollEndEventCondition struct{}

package eventsub

type ConduitShardDisabledEvent struct {
	ConduitID      string      `json:"conduit_id"`      // The ID of the conduit.
	ShardID        string      `json:"shard_id"`        // The ID of the disabled shard.
	Status         string      `json:"status"`          // The new status of the transport.
	Transport      interface{} `json:"transport"`       // The disabled transport.
	Method         string      `json:"method"`          // websocketorwebhook.
	Callback       string      `json:"callback"`        // Optional.
	SessionID      string      `json:"session_id"`      // Optional.
	ConnectedAt    string      `json:"connected_at"`    // Optional.
	DisconnectedAt string      `json:"disconnected_at"` // Optional.
}

package eventsub

type ChannelUpdateEvent struct {
	BroadcasterUserID           string   `json:"broadcaster_user_id"`           // The broadcaster’s user ID.
	BroadcasterUserLogin        string   `json:"broadcaster_user_login"`        // The broadcaster’s user login.
	BroadcasterUserName         string   `json:"broadcaster_user_name"`         // The broadcaster’s user display name.
	Title                       string   `json:"title"`                         // The channel’s stream title.
	Language                    string   `json:"language"`                      // The channel’s broadcast language.
	CategoryID                  string   `json:"category_id"`                   // The channel’s category ID.
	CategoryName                string   `json:"category_name"`                 // The category name.
	ContentClassificationLabels []string `json:"content_classification_labels"` // Array of content classification label IDs currently applied on the Channel.
}

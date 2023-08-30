package eventsub

import (
	"encoding/json"
	"testing"
)

func TestChannelUpdate(t *testing.T) {
	input := `{
        "broadcaster_user_id": "1337",
        "broadcaster_user_login": "cool_user",
        "broadcaster_user_name": "Cool_User",
        "title": "Best Stream Ever",
        "language": "en",
        "category_id": "12453",
        "category_name": "Grand Theft Auto",
        "content_classification_labels": [ "MatureGame" ]
    }`
	expected := ChannelUpdateEvent{
		BroadcasterUserID:           "1337",
		BroadcasterUserLogin:        "cool_user",
		BroadcasterUserName:         "Cool_User",
		Title:                       "Best Stream Ever",
		Language:                    "en",
		CategoryID:                  "12453",
		CategoryName:                "Grand Theft Auto",
		ContentClassificationLabels: []string{"MatureGame"},
	}
	var actual ChannelUpdateEvent
	err := json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	broadcasterEq := expected.BroadcasterUserID == actual.BroadcasterUserID &&
		expected.BroadcasterUserLogin == actual.BroadcasterUserLogin &&
		expected.BroadcasterUserName == actual.BroadcasterUserName
	titleEq := expected.Title == actual.Title
	languageEq := expected.Language == actual.Language
	categoryIDEq := expected.CategoryID == actual.CategoryID
	categoryNameEq := expected.CategoryName == actual.CategoryName
	contentClassificationLabelsEq := len(expected.ContentClassificationLabels) == len(actual.ContentClassificationLabels)

	if contentClassificationLabelsEq {
		for i, v := range expected.ContentClassificationLabels {
			if v != actual.ContentClassificationLabels[i] {
				contentClassificationLabelsEq = false
				break
			}
		}
	}

	if !(broadcasterEq &&
		titleEq &&
		languageEq &&
		categoryIDEq &&
		categoryNameEq &&
		contentClassificationLabelsEq) {
		t.Fatal(eventMismatchErrorMessage(actual, expected))
	}
}

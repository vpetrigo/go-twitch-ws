package eventsub

import (
	"encoding/json"
	"fmt"
	"testing"
)

func eventMismatchErrorMessage(actual, expected any) string {
	return fmt.Sprintf("unmarshalled structure mismatched with expected:\nA: %#v\nE: %#v", actual, expected)
}

func validateInput[Ty comparable](t *testing.T, input string, expected Ty) {
	var u Ty
	err := json.Unmarshal([]byte(input), &u)

	if err != nil {
		t.Fatal(err)
	}

	if u != expected {
		t.Fatal(eventMismatchErrorMessage(u, expected))
	}
}

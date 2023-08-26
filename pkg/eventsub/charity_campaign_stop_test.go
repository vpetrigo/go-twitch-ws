package eventsub

import "testing"

func TestCharityCampaignStop(t *testing.T) {
	input := `{
    "id": "123-abc-456-def",
    "broadcaster_id": "123456",
    "broadcaster_name": "SunnySideUp",
    "broadcaster_login": "sunnysideup",
    "charity_name": "Example name",
    "charity_description": "Example description",
    "charity_logo": "https://abc.cloudfront.net/ppgf/1000/100.png",
    "charity_website": "https://www.example.com",
    "current_amount": {
      "value": 1450000,
      "decimal_places": 2,
      "currency": "USD"
    },
    "target_amount": {
      "value": 1500000,
      "decimal_places": 2,
      "currency": "USD"
    },
    "stopped_at": "2022-07-26T22:00:03.17106713Z"
  }`
	expected := CharityCampaignStopEvent{
		ID:                 "123-abc-456-def",
		BroadcasterID:      "123456",
		BroadcasterName:    "SunnySideUp",
		BroadcasterLogin:   "sunnysideup",
		CharityName:        "Example name",
		CharityDescription: "Example description",
		CharityLogo:        "https://abc.cloudfront.net/ppgf/1000/100.png",
		CharityWebsite:     "https://www.example.com",
		CurrentAmount: CurrentAmount{
			Value:         1450000,
			DecimalPlaces: 2,
			Currency:      "USD",
		},
		TargetAmount: TargetAmount{
			Value:         1500000,
			DecimalPlaces: 2,
			Currency:      "USD",
		},
		StoppedAt: "2022-07-26T22:00:03.17106713Z",
	}

	validateInput(t, input, expected)
}

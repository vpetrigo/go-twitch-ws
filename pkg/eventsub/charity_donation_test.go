package eventsub

import "testing"

func TestCharityDonation(t *testing.T) {
	input := `{
    "id": "a1b2c3-aabb-4455-d1e2f3",
    "campaign_id": "123-abc-456-def",
    "broadcaster_user_id": "123456",
    "broadcaster_user_name": "SunnySideUp",
    "broadcaster_user_login": "sunnysideup",
    "user_id": "654321",
    "user_login": "generoususer1",
    "user_name": "GenerousUser1",
    "charity_name": "Example name",
    "charity_description": "Example description",
    "charity_logo": "https://abc.cloudfront.net/ppgf/1000/100.png",
    "charity_website": "https://www.example.com",
    "amount": {
      "value": 10000,
      "decimal_places": 2,
      "currency": "USD"
    }
  }`
	expected := CharityDonationEvent{
		ID:                   "a1b2c3-aabb-4455-d1e2f3",
		CampaignID:           "123-abc-456-def",
		BroadcasterUserID:    "123456",
		BroadcasterUserName:  "SunnySideUp",
		BroadcasterUserLogin: "sunnysideup",
		UserID:               "654321",
		UserLogin:            "generoususer1",
		UserName:             "GenerousUser1",
		CharityName:          "Example name",
		CharityDescription:   "Example description",
		CharityLogo:          "https://abc.cloudfront.net/ppgf/1000/100.png",
		CharityWebsite:       "https://www.example.com",
		Amount: Amount{
			Value:         10000,
			DecimalPlaces: 2,
			Currency:      "USD",
		},
	}

	validateInput(t, input, expected)
}

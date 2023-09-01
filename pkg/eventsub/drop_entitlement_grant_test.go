package eventsub

import (
	"testing"
)

func TestDropEntitlementGrantUnmarshal(t *testing.T) {
	input := `{
            "id": "bf7c8577-e3e3-4881-a78a-e9446641d45d",
            "data": {
                "organization_id": "9001",
                "category_id": "9002",
                "category_name": "Fortnite",
                "campaign_id": "9003",
                "user_id": "1234",
                "user_name": "Cool_User",
                "user_login": "cool_user",
                "entitlement_id": "fb78259e-fb81-4d1b-8333-34a06ffc24c0",
                "benefit_id": "74c52265-e214-48a6-91b9-23b6014e8041",
                "created_at": "2019-01-28T04:17:53.325Z"
            }
        }`
	expected := DropEntitlementGrantEvent{
		ID: "bf7c8577-e3e3-4881-a78a-e9446641d45d",
		Data: dropEntitlementGrantEventData{
			OrganizationID: "9001",
			CategoryID:     "9002",
			CategoryName:   "Fortnite",
			CampaignID:     "9003",
			UserID:         "1234",
			UserName:       "Cool_User",
			UserLogin:      "cool_user",
			EntitlementID:  "fb78259e-fb81-4d1b-8333-34a06ffc24c0",
			BenefitID:      "74c52265-e214-48a6-91b9-23b6014e8041",
			CreatedAt:      "2019-01-28T04:17:53.325Z",
		},
	}

	validateInput(t, input, expected)
}

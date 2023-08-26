package eventsub

type CharityDonationEvent struct {
	ID                 string `json:"id"`                  // An ID that identifies the donation.
	CampaignID         string `json:"campaign_id"`         // An ID that identifies the charity campaign.
	BroadcasterID      string `json:"broadcaster_id"`      // An ID that identifies the broadcaster that’s running the campaign.
	BroadcasterLogin   string `json:"broadcaster_login"`   // The broadcaster’s login name.
	BroadcasterName    string `json:"broadcaster_name"`    // The broadcaster’s display name.
	UserID             string `json:"user_id"`             // An ID that identifies the user that donated to the campaign.
	UserLogin          string `json:"user_login"`          // The user’s login name.
	UserName           string `json:"user_name"`           // The user’s display name.
	CharityName        string `json:"charity_name"`        // The charity’s name.
	CharityDescription string `json:"charity_description"` // A description of the charity.
	CharityLogo        string `json:"charity_logo"`        // A URL to an image of the charity’s logo.
	CharityWebsite     string `json:"charity_website"`     // A URL to the charity’s website.
	Amount             Amount `json:"amount"`              // An object that contains the amount of money that the user donated.
}

type CharityDonationEventCondition struct{}

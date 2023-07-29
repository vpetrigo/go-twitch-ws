package twitchws

type CharityCampaignStartEvent struct {
	Id                 string      `json:"id"`                  // An ID that identifies the charity campaign.
	BroadcasterId      string      `json:"broadcaster_id"`      // An ID that identifies the broadcaster that’s running the campaign.
	BroadcasterLogin   string      `json:"broadcaster_login"`   // The broadcaster’s login name.
	BroadcasterName    string      `json:"broadcaster_name"`    // The broadcaster’s display name.
	CharityName        string      `json:"charity_name"`        // The charity’s name.
	CharityDescription string      `json:"charity_description"` // A description of the charity.
	CharityLogo        string      `json:"charity_logo"`        // A URL to an image of the charity’s logo.
	CharityWebsite     string      `json:"charity_website"`     // A URL to the charity’s website.
	CurrentAmount      interface{} `json:"current_amount"`      // An object that contains the current amount of donations that the campaign has received.
	Value              int         `json:"value"`               // The monetary amount.
	DecimalPlaces      int         `json:"decimal_places"`      // The number of decimal places used by the currency.
	Currency           string      `json:"currency"`            // The ISO-4217 three-letter currency code that identifies the type of currency in value.
	TargetAmount       interface{} `json:"target_amount"`       // An object that contains the campaign’s target fundraising goal.
	Value              int         `json:"value"`               // The monetary amount.
	DecimalPlaces      int         `json:"decimal_places"`      // The number of decimal places used by the currency.
	Currency           string      `json:"currency"`            // The ISO-4217 three-letter currency code that identifies the type of currency in value.
	StartedAt          string      `json:"started_at"`          // The UTC timestamp (in RFC3339 format) of when the broadcaster started the campaign.
}

type CharityCampaignStartEventCondition struct{}

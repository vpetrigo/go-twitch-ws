package twitchws

type CharityCampaignStopEvent struct {
	Id                 string                                `json:"id"`                  // An ID that identifies the charity campaign.
	BroadcasterId      string                                `json:"broadcaster_id"`      // An ID that identifies the broadcaster that ran the campaign.
	BroadcasterLogin   string                                `json:"broadcaster_login"`   // The broadcaster’s login name.
	BroadcasterName    string                                `json:"broadcaster_name"`    // The broadcaster’s display name.
	CharityName        string                                `json:"charity_name"`        // The charity’s name.
	CharityDescription string                                `json:"charity_description"` // A description of the charity.
	CharityLogo        string                                `json:"charity_logo"`        // A URL to an image of the charity’s logo.
	CharityWebsite     string                                `json:"charity_website"`     // A URL to the charity’s website.
	CurrentAmount      charityCampaignStopEventCurrentAmount `json:"current_amount"`      // An object that contains the final amount of donations that the campaign received.
	TargetAmount       charityCampaignStopEventTargetAmount  `json:"target_amount"`       // An object that contains the campaign’s target fundraising goal.
	StoppedAt          string                                `json:"stopped_at"`          // The UTC timestamp (in RFC3339 format) of when the broadcaster stopped the campaign.
}

type charityCampaignStopEventCurrentAmount struct {
	Value         int    `json:"value"`          // The monetary amount.
	DecimalPlaces int    `json:"decimal_places"` // The number of decimal places used by the currency.
	Currency      string `json:"currency"`       // The ISO-4217 three-letter currency code that identifies the type of currency in value.
}

type charityCampaignStopEventTargetAmount struct {
	Value         int    `json:"value"`          // The monetary amount.
	DecimalPlaces int    `json:"decimal_places"` // The number of decimal places used by the currency.
	Currency      string `json:"currency"`       // The ISO-4217 three-letter currency code that identifies the type of currency in value.
}

type CharityCampaignStopEventCondition struct{}

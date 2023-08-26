package eventsub

type CharityCampaignProgressEvent struct {
	ID                 string        `json:"id"`                  // An ID that identifies the charity campaign.
	BroadcasterID      string        `json:"broadcaster_id"`      // An ID that identifies the broadcaster that’s running the campaign.
	BroadcasterLogin   string        `json:"broadcaster_login"`   // The broadcaster’s login name.
	BroadcasterName    string        `json:"broadcaster_name"`    // The broadcaster’s display name.
	CharityName        string        `json:"charity_name"`        // The charity’s name.
	CharityDescription string        `json:"charity_description"` // A description of the charity.
	CharityLogo        string        `json:"charity_logo"`        // A URL to an image of the charity’s logo.
	CharityWebsite     string        `json:"charity_website"`     // A URL to the charity’s website.
	CurrentAmount      CurrentAmount `json:"current_amount"`      // An object that contains the current amount of donations that the campaign has received.
	TargetAmount       TargetAmount  `json:"target_amount"`       // An object that contains the campaign’s target fundraising goal.
}

type CharityCampaignProgressEventCondition struct{}

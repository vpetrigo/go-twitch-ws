package eventsub

type AutomodSettingsUpdateEvent struct {
	BroadcasterUserID       string      `json:"broadcaster_user_id"`        // The ID of the broadcaster specified in the request.
	BroadcasterUserLogin    string      `json:"broadcaster_user_login"`     // The login of the broadcaster specified in the request.
	BroadcasterUserName     string      `json:"broadcaster_user_name"`      // The user name of the broadcaster specified in the request.
	ModeratorUserID         string      `json:"moderator_user_id"`          // The ID of the moderator who changed the channel settings.
	ModeratorUserLogin      string      `json:"moderator_user_login"`       // The moderator’s login.
	ModeratorUserName       string      `json:"moderator_user_name"`        // The moderator’s user name.
	Bullying                int         `json:"bullying"`                   // The Automod level for hostility involving name calling or insults.
	OverallLevel            interface{} `json:"overall_level"`              // The default AutoMod level for the broadcaster.
	Disability              int         `json:"disability"`                 // The Automod level for discrimination against disability.
	RaceEthnicityOrReligion int         `json:"race_ethnicity_or_religion"` // The Automod level for racial discrimination.
	Misogyny                int         `json:"misogyny"`                   // The Automod level for discrimination against women.
	SexualitySexOrGender    int         `json:"sexuality_sex_or_gender"`    // The AutoMod level for discrimination based on sexuality, sex, or gender.
	Aggression              int         `json:"aggression"`                 // The Automod level for hostility involving aggression.
	SexBasedTerms           int         `json:"sex_based_terms"`            // The Automod level for sexual content.
	Swearing                int         `json:"swearing"`                   // The Automod level for profanity.
}

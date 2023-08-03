package eventsub

type DropEntitlementGrantEvent struct {
	ID   string                        `json:"id"`   // Individual event ID, as assigned by EventSub.
	Data dropEntitlementGrantEventData `json:"data"` // Entitlement object.
}

type dropEntitlementGrantEventData struct {
	OrganizationID string `json:"organization_id"` // The ID of the organization that owns the game that has Drops enabled.
	CategoryID     string `json:"category_id"`     // Twitch category ID of the game that was being played when this benefit was entitled.
	CategoryName   string `json:"category_name"`   // The category name.
	CampaignID     string `json:"campaign_id"`     // The campaign this entitlement is associated with.
	UserID         string `json:"user_id"`         // Twitch user ID of the user who was granted the entitlement.
	UserName       string `json:"user_name"`       // The user display name of the user who was granted the entitlement.
	UserLogin      string `json:"user_login"`      // The user login of the user who was granted the entitlement.
	EntitlementID  string `json:"entitlement_id"`  // Unique identifier of the entitlement.
	BenefitID      string `json:"benefit_id"`      // Identifier of the Benefit.
	CreatedAt      string `json:"created_at"`      // UTC timestamp in ISO format when this entitlement was granted on Twitch.
}

type DropEntitlementGrantEventCondition struct{}

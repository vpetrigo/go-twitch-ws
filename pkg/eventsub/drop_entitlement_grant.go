package eventsub

type DropEntitlementGrantEvent struct {
	ID   string                        `json:"id"`   // Individual event ID, as assigned by EventSub.
	Data dropEntitlementGrantEventData `json:"data"` // Entitlement object.
}

type DropEntitlementGrantEventCondition struct{}

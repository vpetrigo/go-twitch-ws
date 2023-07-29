package twitchws

type DropEntitlementGrantEvent struct {
	Id   string      `json:"id"`   // Individual event ID, as assigned by EventSub.
	Data interface{} `json:"data"` // Entitlement object.
}

type DropEntitlementGrantEventCondition struct{}

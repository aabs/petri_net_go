package persistence

import (
	"aabs/petri_net_go/core"
	"time"
)

// payload types:
// - on new asset onboarded
// - on new target request
// - on telemetry summary (evidence of activity)
// - on target status observation
// - on target request dispatched
// - on new dispatch engine online/offline/etc
// - on new trading agent built/registered/running/etc

type Event struct {
	EventId            string      `json:"eventId"`
	Version            string      `json:"version"`
	EventType          string      `json:"eventType"`
	GeneratedTimestamp int64       `json:"generatedTimestamp"` // result of calling time.Now().UnixNano()
	Payload            interface{} `json:"payload"`
}

type AssetFoundDuringOnboardingEvent struct {
	RecipientId string `json:"recipient_id"`
	AssetId     string `json:"asset_id"`
}

type NewTargetRequestEvent struct {
	// get this from common/targets.go when integrated into encore
	TargetID          string    `json:"target_id"`
	Start             time.Time `json:"start"`
	End               time.Time `json:"end"`
	Mode              string    `json:"mode"`
	Priority          int64     `json:"priority"`
	SetPoint          float64   `json:"setpoint"`
	TechVerification  bool      `json:"techVerification"`
	IgnoreEligibility bool      `json:"ignoreEligibility"`
	UseOSPC           bool      `json:"useOSPC"`
}

type NewPetriNetDefinitionEvent struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Places      []core.Place      `json:"places"`
	Transitions []core.Transition `json:"transitions"`
	InArcs      []core.Arc        `json:"inArcs"`
	OutArcs     []core.Arc        `json:"outArcs"`
	InhibArcs   []core.Arc        `json:"inhibArcs"`
}

type Marking struct {
	PetriNetName    string `json:"petriNetName"`
	PetriNetVersion string `json:"petriNetVersion"`
	Tokens          []int  `json:"tokens"`
}

type FiringVector struct {
	PetriNetName    string `json:"petriNetName"`
	PetriNetVersion string `json:"petriNetVersion"`
	Transitions     []bool `json:"transitions"`
}

type NewMarkingReachedEvent struct {
	AssetId        string  `json:"asset_id"`
	MarkingReached Marking `json:"markingReached"`
}

type NewFiringVectorEvent struct {
	AssetId          string       `json:"asset_id"`
	FiredTransitions FiringVector `json:"firedTransitions"`
}

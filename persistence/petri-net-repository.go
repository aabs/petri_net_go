package persistence

import (
	"aabs/petri_net_go/core"
)

type PetriNetRepository interface {
	InsertDefinition(net *core.PetriNet) error
	GetDefinition(name string, version string) (*core.PetriNet, error)
	UpdateDefinition(net *core.PetriNet) error
	DeleteDefinition(name string, version string) error

	InsertMarking(net *core.Marking) error
	GetMarking(instanceId string, name string) (*core.Marking, error)
	UpdateMarking(net *core.Marking) error
	DeleteMarking(instanceId string, name string) error
}

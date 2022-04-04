package persistence

import (
	"aabs/petri_net_go/core"
	"errors"
)

type TemporaryInMemoryPetriNetRepository struct {
	definitions map[string]*core.PetriNet
	markings    map[string]*core.Marking
}

func (repo *TemporaryInMemoryPetriNetRepository) InsertDefinition(net *core.PetriNet) error {
	if repo == nil {
		return errors.New("repo is nil")
	}
	if net == nil {
		return errors.New("PetriNet is nil")
	}
	repo.definitions[net.Name] = net
	return nil
}

func (repo *TemporaryInMemoryPetriNetRepository) GetDefinition(name string, version string) (*core.PetriNet, error) {
	result := repo.definitions[name]
	return result.Clone(), nil
}

func (repo *TemporaryInMemoryPetriNetRepository) UpdateDefinition(net *core.PetriNet) error {
	repo.definitions[net.Name] = net
	return nil
}

func (repo *TemporaryInMemoryPetriNetRepository) DeleteDefinition(name string, version string) error {
	delete(repo.definitions, name)
	return nil
}

func (repo *TemporaryInMemoryPetriNetRepository) InsertMarking(net *core.Marking) error {
	repo.markings[net.InstanceId] = net
	return nil
}

func (repo *TemporaryInMemoryPetriNetRepository) GetMarking(instanceId string, name string) (*core.Marking, error) {
	return repo.markings[instanceId], nil
}

func (repo *TemporaryInMemoryPetriNetRepository) UpdateMarking(net *core.Marking) error {
	repo.markings[net.InstanceId] = net
	return nil
}

func (repo *TemporaryInMemoryPetriNetRepository) DeleteMarking(instanceId string, name string) error {
	delete(repo.markings, instanceId)
	return nil
}

func NewTemporaryInMemoryPetriNetRepository() *TemporaryInMemoryPetriNetRepository {
	return &TemporaryInMemoryPetriNetRepository{
		definitions: make(map[string]*core.PetriNet),
		markings:    make(map[string]*core.Marking),
	}
}

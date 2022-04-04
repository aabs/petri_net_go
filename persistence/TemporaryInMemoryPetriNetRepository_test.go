package persistence

import (
	"aabs/petri_net_go/core"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemporaryInMemoryPetriNetRepository_InsertDefinition(t *testing.T) {
	type args struct {
		net *core.PetriNet
	}
	tests := []struct {
		name    string
		repo    *TemporaryInMemoryPetriNetRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "can store a petrinet",
			args: args{
				net: CreateTestDefinition(),
			},
			wantErr: false,
			repo:    NewTemporaryInMemoryPetriNetRepository(),
		},
		{
			name: "will refuse store a nil petrinet",
			args: args{
				net: nil,
			},
			wantErr: true,
			repo:    NewTemporaryInMemoryPetriNetRepository(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.repo.InsertDefinition(tt.args.net); (err != nil) != tt.wantErr {
				t.Errorf("TemporaryInMemoryPetriNetRepository.CreateDefinition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTemporaryInMemoryPetriNetRepository_GetDefinition(t *testing.T) {
	sut := NewTemporaryInMemoryPetriNetRepository()
	net := CreateTestDefinition()
	sut.InsertDefinition(net)

	// act
	result, err := sut.GetDefinition(net.Name, net.Version)

	// assert
	assert.Nil(t, err)
	CheckNetEquality(t, result, net, true)
}

func CheckNetEquality(t *testing.T, net *core.PetriNet, net2 *core.PetriNet, shouldBeEqual bool) {
	areEqual := true
	areEqual = areEqual && reflect.DeepEqual(net.Name, net2.Name)
	areEqual = areEqual && reflect.DeepEqual(net.Version, net2.Version)
	areEqual = areEqual && reflect.DeepEqual(net.InhibitoryInputIncidence, net2.InhibitoryInputIncidence)
	areEqual = areEqual && reflect.DeepEqual(net.InputIncidence, net2.InputIncidence)
	areEqual = areEqual && reflect.DeepEqual(net.OutputIncidence, net2.OutputIncidence)
	assert.Equal(t, shouldBeEqual, areEqual)
}

func TestTemporaryInMemoryPetriNetRepository_UpdateDefinition(t *testing.T) {
	sut := NewTemporaryInMemoryPetriNetRepository()
	net := CreateTestDefinition()
	sut.InsertDefinition(net)

	// act
	result, err := sut.GetDefinition(net.Name, net.Version)

	// assert
	assert.Nil(t, err)
	CheckNetEquality(t, result, net, true)

	net.Name = "some new name"
	net.Version = "1.0.1"
	CheckNetEquality(t, result, net, false)

	err = sut.UpdateDefinition(net)
	assert.Nil(t, err)

	result, err = sut.GetDefinition(net.Name, net.Version)
	assert.Nil(t, err)
	CheckNetEquality(t, result, net, true)

}

func TestTemporaryInMemoryPetriNetRepository_DeleteDefinition(t *testing.T) {
	sut := NewTemporaryInMemoryPetriNetRepository()
	net := CreateTestDefinition()
	sut.InsertDefinition(net)

	// act
	result, err := sut.GetDefinition(net.Name, net.Version)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestTemporaryInMemoryPetriNetRepository_InsertMarking(t *testing.T) {
	type args struct {
		net *core.Marking
	}
	tests := []struct {
		name    string
		repo    *TemporaryInMemoryPetriNetRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "can store a marking",
			args: args{
				net: CreateTestMarking(),
			},
			wantErr: false,
			repo:    NewTemporaryInMemoryPetriNetRepository(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.repo.InsertMarking(tt.args.net); (err != nil) != tt.wantErr {
				t.Errorf("TemporaryInMemoryPetriNetRepository.CreateMarking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTemporaryInMemoryPetriNetRepository_GetMarking(t *testing.T) {
	type args struct {
		instanceId string
		name       string
	}
	tests := []struct {
		name    string
		repo    *TemporaryInMemoryPetriNetRepository
		args    args
		want    *core.Marking
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "can store a marking",
			args: args{
				instanceId: "testId",
				name:       "PT",
			},
			wantErr: false,
			repo:    NewTemporaryInMemoryPetriNetRepository(),
			want:    CreateTestMarking(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.InsertMarking(tt.want)
			got, err := tt.repo.GetMarking(tt.args.instanceId, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TemporaryInMemoryPetriNetRepository.GetMarking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TemporaryInMemoryPetriNetRepository.GetMarking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemporaryInMemoryPetriNetRepository_UpdateMarking(t *testing.T) {
	type args struct {
		marking *core.Marking
	}
	tests := []struct {
		name    string
		repo    *TemporaryInMemoryPetriNetRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "can update a marking",
			args: args{
				marking: CreateTestMarking(),
			},
			wantErr: false,
			repo:    NewTemporaryInMemoryPetriNetRepository(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.InsertMarking(tt.args.marking)
			got, err := tt.repo.GetMarking(tt.args.marking.InstanceId, tt.args.marking.DefinitionName)
			if !reflect.DeepEqual(got, tt.args.marking) {
				t.Errorf("TemporaryInMemoryPetriNetRepository.GetMarking() = %v, want %v", got, tt.args.marking)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("TemporaryInMemoryPetriNetRepository.UpdateMarking() error = %v, wantErr %v", err, tt.wantErr)
			}
			m_1, err2 := got.SetMarking(int(p2), 12.3)
			if err2 = tt.repo.UpdateMarking(m_1); (err2 != nil) != tt.wantErr {
				t.Errorf("TemporaryInMemoryPetriNetRepository.UpdateMarking() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(m_1, tt.args.marking) {
				t.Errorf("TemporaryInMemoryPetriNetRepository.GetMarking() = %v, want %v", got, tt.args.marking)
			}
		})
	}
}

func TestTemporaryInMemoryPetriNetRepository_DeleteMarking(t *testing.T) {
	type args struct {
		instanceId string
		name       string
	}
	tests := []struct {
		name    string
		repo    *TemporaryInMemoryPetriNetRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.repo.DeleteMarking(tt.args.instanceId, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("TemporaryInMemoryPetriNetRepository.DeleteMarking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewTemporaryInMemoryPetriNetRepository(t *testing.T) {
	sut := NewTemporaryInMemoryPetriNetRepository()
	assert.NotNil(t, sut)
}

const (
	p1 core.PlaceId = iota
	p2
	p3
	p4
)

const (
	t1 core.TransitionId = iota
	t2
	t3
)

// HELPER FUNCTIONS
func CreateTestDefinition() *core.PetriNet {
	result, _ := core.CreatePetriNet().
		Called("PT").
		WithPlaces(map[core.PlaceId]string{p1: "p1", p2: "p2", p3: "p3", p4: "p4"}).
		WithTransitions(map[core.TransitionId]string{t1: "t1", t2: "t2", t3: "t3"}).
		WithArcsIntoPlaces(map[core.PlaceId][]core.TransitionId{p1: {t2, t3}, p2: {t1}, p3: {t1}, p4: {t3, t3}}).
		WithArcsOutOfPlaces(map[core.PlaceId][]core.TransitionId{p1: {t1}, p2: {t2}, p3: {t3}, p4: {t2, t2}}).
		Build()
	return result
}

func CreateTestMarking() *core.Marking {
	result := core.CreateMarking(4, []int{2, 0, 1, 0})
	result.InstanceId = "testId"
	result.DefinitionName = "PT"
	result.DefinitionVersion = "1.2.3"
	return &result
}

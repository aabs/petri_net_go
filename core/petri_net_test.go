package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	P1 PlaceId = iota
	P2
	P3
	P4
)

const (
	T1 TransitionId = iota
	T2
	T3
)

/*func Test_petriNet_NewPetriNet(t *testing.T) {
	sut, err := NewPetriNet(5, 5)

	if err != nil {
		t.Error(sut)
	}
}
*/
func Test_PlaceBuilder_Create(t *testing.T) {
	sut := CreatePlace().WithId(13)

	if sut.Id != 13 {
		t.Error(sut)
	}
}
func Test_PlaceBuilder_Called(t *testing.T) {
	sut := CreatePlace().
		Called("Test")

	if sut.Name != "Test" {
		t.Error(sut)
	}
}
func Test_TransitionBuilder_Create(t *testing.T) {
	sut := CreateTransition()

	if sut.Id != 0 {
		t.Error(sut)
	}
}
func Test_TransitionBuilder_Called(t *testing.T) {
	sut := CreateTransition().
		Called("Test")

	if sut.Name != "Test" {
		t.Error(sut)
	}
}

func Test_PetriNetBuilder_Create(t *testing.T) {
	sut := CreatePetriNet().
		Called("PT")

	if sut.Name != "PT" {
		t.Error(sut)
	}
}

func CreateTestNet() (*PetriNet, error) {
	return CreatePetriNet().
		Called("PT").
		WithPlaces(map[PlaceId]string{P1: "P1", P2: "P2", P3: "P3", P4: "P4"}).
		WithTransitions(map[TransitionId]string{T1: "T1", T2: "T2", T3: "T3"}).
		WithArcsIntoPlaces(map[PlaceId][]TransitionId{P1: {T2, T3}, P2: {T1}, P3: {T1}, P4: {T3, T3}}).
		WithArcsOutOfPlaces(map[PlaceId][]TransitionId{P1: {T1}, P2: {T2}, P3: {T3}, P4: {T2, T2}}).
		Build()
}

func testErr(e error, t *testing.T) {
	if e != nil {
		t.Error(e)
	}
}
func testBool(e bool, t *testing.T) {
	if !e {
		t.Error(e)
	}
}
func Test_PetriNetBuilder_FullCreate(t *testing.T) {
	sut, err := CreateTestNet()

	testErr(err, t)
	testBool(sut.Name == "PT", t)
}

func Test_PetriNet_StateEqn_1(t *testing.T) {
	net, err := CreateTestNet()
	testErr(err, t)
	marking := CreateMarking(4, []int{2, 0, 1, 0})

	firingList, err := net.GetEligibleFiringList(marking)
	testErr(err, t)
	assert.Equal(t, firingList.Len(), 3, "")
	assert.Equal(t, firingList.At(0, 0), 1.0, "")
	assert.Equal(t, firingList.At(1, 0), 0.0, "")
	assert.Equal(t, firingList.At(2, 0), 1.0, "")
}
func Test_PetriNet_GetEligibleFiringList(t *testing.T) {
	net, err := CreateTestNet()
	testErr(err, t)
	marking := CreateMarking(4, []int{2, 0, 1, 0})

	firingList, err := net.GetEligibleFiringList(marking)
	testErr(err, t)
	assert.Equal(t, firingList.Len(), 3, "")
	assert.Equal(t, firingList.At(0, 0), 1.0, "")
	assert.Equal(t, firingList.At(1, 0), 0.0, "")
	assert.Equal(t, firingList.At(2, 0), 1.0, "")
	marking2 := CreateMarking(4, []int{0, 1, 0, 2})

	firingList2, _ := net.GetEligibleFiringList(marking2)
	testErr(err, t)
	assert.Equal(t, firingList2.Len(), 3, "")
	assert.Equal(t, firingList2.At(0, 0), 0.0, "")
	assert.Equal(t, firingList2.At(1, 0), 1.0, "")
	assert.Equal(t, firingList2.At(2, 0), 0.0, "")
}

func Test_PetriNet_StateEqn_2(t *testing.T) {
	// arrange
	net, _ := CreateTestNet()
	marking := CreateMarking(4, []int{2, 0, 1, 0})
	firingList, _ := net.GetEligibleFiringList(marking)
	chosenTransition, err := net.ChooseTransitionFromEligibleFiringList(firingList)
	testErr(err, t)
	assert.True(t, chosenTransition.At(0, 0) == 1.0 || chosenTransition.At(2, 0) == 1.0, "only T1 or T3 should have been chosen")
	assert.True(t, chosenTransition.At(1, 0) == 0.0, "T2 should never have been chosen")
}

func Test_PetriNet_InvokesTransitionHandler(t *testing.T) {
	testCounter := 0
	sut, err := CreatePetriNet().
		Called("PT").
		WithPlaces(map[PlaceId]string{P1: "P1", P2: "P2"}).
		WithTransitions(map[TransitionId]string{T1: "T1"}).
		WithTransitionHandler(T1, func(tid TransitionId, m_0 *Marking, m_1 *Marking) { testCounter++ }).
		WithArcsIntoPlaces(map[PlaceId][]TransitionId{P2: {T1}}).
		WithArcsOutOfPlaces(map[PlaceId][]TransitionId{P1: {T1}}).
		Build()
	testErr(err, t)
	m := CreateMarking(2, []int{1, 0})

	u, err2 := sut.GetEligibleFiringList(m)
	assert.Nil(t, err2)
	assert.NotNil(t, u)
	assert.Equal(t, testCounter, 0)

	var m_1 *Marking
	m_1, err = sut.Fire(&m, u)
	assert.Nil(t, err)
	assert.NotNil(t, m_1)
	assert.Equal(t, testCounter, 1)
}

func Test_PetriNet_InhibitorCausesTransitionToBeIgnored(t *testing.T) {
	sut, _ := CreatePetriNet().
		Called("PT").
		WithPlaces(map[PlaceId]string{P1: "P1", P2: "P2", P3: "P3", P4: "PGuard"}).
		WithTransitions(map[TransitionId]string{T1: "T1", T2: "T2"}).
		WithArcsIntoPlaces(map[PlaceId][]TransitionId{P2: {T1}, P3: {T2}}).
		WithArcsOutOfPlaces(map[PlaceId][]TransitionId{P1: {T1, T2}}).
		WithInhibitorArc(P4, T2).
		Build()
	m := CreateMarking(4, []int{1, 0, 0, 1})

	u, err2 := sut.GetEligibleFiringList(m)
	testErr(err2, t)
	assert.Equal(t, u.Len(), 2, "only two transitions were defined")
	assert.Equal(t, u.At(0, 0), 1.0, "this should have been enabled")
	assert.Equal(t, u.At(1, 0), 0.0, "this transition should have been inhibited")
}

func Test_PetriNet_InhibitorCausesTransitionToBeIgnoredOnlyWhenTokensMeetThreshold1(t *testing.T) {
	sut, _ := CreatePetriNet().
		Called("PT").
		WithPlaces(map[PlaceId]string{P1: "P1", P2: "P2", P3: "P3", P4: "PGuard"}).
		WithTransitions(map[TransitionId]string{T1: "T1", T2: "T2"}).
		WithArcsIntoPlaces(map[PlaceId][]TransitionId{P2: {T1}, P3: {T2}}).
		WithArcsOutOfPlaces(map[PlaceId][]TransitionId{P1: {T1, T2}}).
		WithInhibitorArcs(map[PlaceId][]TransitionId{P4: {T2, T2}}). // i.e. a weight of two between P4 and T2
		Build()
	m := CreateMarking(4, []int{1, 0, 0, 1})

	u, err2 := sut.GetEligibleFiringList(m)
	testErr(err2, t)
	assert.Equal(t, u.Len(), 2, "only two transitions were defined")
	assert.Equal(t, u.At(0, 0), 1.0, "this should have been enabled")
	assert.Equal(t, u.At(1, 0), 1.0, "this transition should have been inhibited")
}

func Test_PetriNet_InhibitorCausesTransitionToBeIgnoredOnlyWhenTokensMeetThreshold2(t *testing.T) {
	sut, _ := CreatePetriNet().
		Called("PT").
		WithPlaces(map[PlaceId]string{P1: "P1", P2: "P2", P3: "P3", P4: "PGuard"}).
		WithTransitions(map[TransitionId]string{T1: "T1", T2: "T2"}).
		WithArcsIntoPlaces(map[PlaceId][]TransitionId{P2: {T1}, P3: {T2}}).
		WithArcsOutOfPlaces(map[PlaceId][]TransitionId{P1: {T1, T2}}).
		WithInhibitorArcs(map[PlaceId][]TransitionId{P4: {T2, T2}}). // i.e. a weight of two between P4 and T2
		Build()
	m := CreateMarking(4, []int{1, 0, 0, 2})

	u, err2 := sut.GetEligibleFiringList(m)
	testErr(err2, t)
	assert.Equal(t, u.Len(), 2, "only two transitions were defined")
	assert.Equal(t, u.At(0, 0), 1.0, "this should have been enabled")
	assert.Equal(t, u.At(1, 0), 0.0, "this transition should have been inhibited")
}

func Test_PetriNet_InhibitorCausesTransitionToBeIgnoredOnlyWhenTokensMeetThreshold3(t *testing.T) {
	sut, _ := CreatePetriNet().
		Called("PT").
		WithPlaces(map[PlaceId]string{P1: "P1", P2: "P2", P3: "P3", P4: "PGuard"}).
		WithTransitions(map[TransitionId]string{T1: "T1", T2: "T2"}).
		WithArcsIntoPlaces(map[PlaceId][]TransitionId{P2: {T1}, P3: {T2}}).
		WithArcsOutOfPlaces(map[PlaceId][]TransitionId{P1: {T1, T2}}).
		WithInhibitorArcs(map[PlaceId][]TransitionId{P4: {T2, T2}}). // i.e. a weight of two between P4 and T2
		Build()
	m := CreateMarking(4, []int{1, 0, 0, 345})

	u, err2 := sut.GetEligibleFiringList(m)
	testErr(err2, t)
	assert.Equal(t, u.Len(), 2, "only two transitions were defined")
	assert.Equal(t, u.At(0, 0), 1.0, "this should have been enabled")
	assert.Equal(t, u.At(1, 0), 0.0, "this transition should have been inhibited")
}

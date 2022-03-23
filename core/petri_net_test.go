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

	if sut.Id != 1 {
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
		WithPlace(*CreatePlace().WithId(P1).Called("P1").Build()).
		WithPlace(*CreatePlace().WithId(P2).Called("P2").Build()).
		WithPlace(*CreatePlace().WithId(P3).Called("P3").Build()).
		WithPlace(*CreatePlace().WithId(P4).Called("P4").Build()).
		WithTransition(*CreateTransition().WithId(T1).Called("T1").Build()).
		WithTransition(*CreateTransition().WithId(T2).Called("T2").Build()).
		WithTransition(*CreateTransition().WithId(T3).Called("T3").Build()).
		WithInArcs(map[PlaceId][]TransitionId{P1: {T2, T3}, P2: {T1}, P3: {T1}, P4: {T3, T3}}).
		WithOutArcs(map[PlaceId][]TransitionId{P1: {T1}, P2: {T2}, P3: {T3}, P4: {T2, T2}}).
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
	marking := CreateMarking(4, []int{2, 0, 1, 0, })

	firingList, err := net.GetFiringList(marking)
	testErr(err, t)
	testBool(firingList.Len() == 3, t)
	testBool(firingList.At(0, 0) == 0, t)
	testBool(firingList.At(1, 0) == 0, t)
	testBool(firingList.At(2, 0) == 1, t)
}

func Test_PetriNet_StateEqn_2(t *testing.T) {
	// arrange
	net, _ := CreateTestNet()
	marking := CreateMarking(4, []int{2, 0, 1, 0, })
	firingList, _ := net.GetFiringList(marking)
	assert.Equal(t, 1.0, marking.Places.AtVec(0), "initial marking")
	assert.Equal(t, 0.0, marking.Places.AtVec(1), "initial marking")

	// act
	newMarking, err := net.Fire(marking, firingList)

	// assert
	testErr(err, t)
	assert.Equal(t, 0.0, newMarking.Places.AtVec(0), "marking should have moved")
	assert.Equal(t, 1.0, newMarking.Places.AtVec(1), "marking should have moved")
}

package core

import (
	"testing"
)

/*func Test_petriNet_NewPetriNet(t *testing.T) {
	sut, err := NewPetriNet(5, 5)

	if err != nil {
		t.Error(sut)
	}
}
*/
func Test_PlaceBuilder_Create(t *testing.T) {
	sut := CreatePlace()

	if sut.Id != 1 {
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
		WithPlace(*CreatePlace().Called("P1").Build()).
		WithPlace(*CreatePlace().Called("P2").Build()).
		WithTransition(*CreateTransition().Called("T1").Build()).
		WithInArc(*NewPair(1, 1)).
		WithOutArc(*NewPair(1, 2)).
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
	marking := CreateMarking(2, []int{0, 0})

	firingList, err := net.GetFiringList(marking)
	testErr(err, t)
	testBool(firingList.Len() == 1, t)
	testBool(firingList.At(0, 0) == 1, t)
}

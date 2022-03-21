package core

import (
	"testing"
)

func Test_petriNet_NewPetriNet(t *testing.T) {
	sut, err := NewPetriNet(5, 5)

	if err != nil {
		t.Error(sut)
	}
}

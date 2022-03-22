// Package core implements basic definitions and representations of Petri nets.

package core

import (
	"gonum.org/v1/gonum/mat"
)

// PetriNet holds the specification of the petri net.
type PetriNet struct {
	Name                     string
	Version                  string
	InputIncidence           mat.Matrix
	OutputIncidence          mat.Matrix
	InhibitoryInputIncidence mat.Matrix
	PlaceNames               map[int]string
	TransitionNames          map[int]string
}

func (net *PetriNet) GetFiringList(marking Marking) (*mat.VecDense, error) {
	number_of_transitions := len(net.TransitionNames)
	firings := convertTo64([]int{1})
	return mat.NewVecDense(number_of_transitions, firings), nil
}

type Marking struct {
	Places mat.Vector
}

func convertTo64(ar []int) []float64 {
	newar := make([]float64, len(ar))
	var v int
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}


func CreateMarking(size int, markings []int) Marking {
	places := mat.NewVecDense(size, convertTo64(markings))
	return Marking{
		Places: places,
	}
}

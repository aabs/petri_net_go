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

type PlaceId int
type TransitionId int

type Arc struct {
	Place      PlaceId
	Transition TransitionId
}

func (net *PetriNet) GetFiringList(marking Marking) (*mat.VecDense, error) {
	number_of_transitions := len(net.TransitionNames)
	firings := ConvertToFloat64([]int{0, 0, 1})
	return mat.NewVecDense(number_of_transitions, firings), nil
}

func (net *PetriNet) Fire(m_0 Marking, firingList *mat.VecDense) (*Marking, error) {
	var A mat.Dense
	A.Sub(net.InputIncidence, net.OutputIncidence)
	//AT := A.T()
	var d mat.VecDense
	d.MulVec(&A, firingList)
	var m_1 mat.VecDense
	m_1.AddVec(m_0.Places, &d)
	//newMarking := m_0 + (net.InputIncidence - net.OutputIncidence ) * firingList
	return &Marking{
		Places: &m_1,
	}, nil
}

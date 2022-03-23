// Package core implements basic definitions and representations of Petri nets.

package core

import (
	"math/rand"

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

func (net *PetriNet) GetEligibleFiringList(marking Marking) (*mat.VecDense, error) {
	numPlaces := marking.Places.Len()
	numTransitions := len(net.TransitionNames)
	firingList := make([]float64, numTransitions, numTransitions)

	for t := 0; t < numTransitions; t++ {
		enabled := true
		for p := 0; p < numPlaces; p++ {
			a_pt := net.OutputIncidence.At(p, t) // how many tokens does t require from p
			m_p := marking.Places.At(p, 0)       // how many tokens does p have
			enabled = enabled && (a_pt <= m_p)   // still possible if the marking for p has more tokens than t needs from p
		}
		boolAsFloat := 0.0
		if enabled {
			boolAsFloat = 1.0
		}
		firingList[t] = boolAsFloat
	}

	return mat.NewVecDense(numTransitions, firingList), nil
}

func (net *PetriNet) ChooseTransitionFromEligibleFiringList(fullFiringList *mat.VecDense) (*mat.VecDense, error) {
	numTransitions := len(net.TransitionNames)
	firingList := make([]float64, numTransitions, numTransitions)

	// before returning the firing list we need to pick one of the enabled transitions as the only one to be fired
	// first if only one transition is enabled, we return the firing list as is
	// if more than one transition is enabled, we return the firing list with a randomly chosen (free choice) enabled transition set to 1.0
	// and all the other enabled transitions set to 0.0.
	// if no transition is enabled, we have deadlock, but this function can do nothing about it.

	enabledCount := 0
	for i := 0; i < numTransitions; i++ {
		enabledCount += int(fullFiringList.AtVec(i))
	}
	if enabledCount == 1 {
		return mat.NewVecDense(numTransitions, firingList), nil
	} else if enabledCount > 1 {
		// randomly choose one of the enabled transitions
		enabledTransitionsToDiscard := rand.Intn(enabledCount+1)

		for i := 0; i < numTransitions; i++ {
			if fullFiringList.AtVec(i) == 1.0 {
				enabledTransitionsToDiscard--
				if enabledTransitionsToDiscard == 0 {
					firingList[i] = 1.0
				} else {
					firingList[i] = 0.0
				}
			}
		}
	}

	return mat.NewVecDense(numTransitions, firingList), nil
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

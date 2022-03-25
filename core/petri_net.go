// Package core implements basic definitions and representations of Petri nets.

package core

import (
	"errors"
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
	TransitionHandlers       map[TransitionId]func(TransitionId, *Marking, *Marking)
}

type PlaceId int
type TransitionId int

type Arc struct {
	Place      PlaceId
	Transition TransitionId
}

func ComputeVectorOfInhibitedTransitions(m Marking, p *PetriNet) *mat.VecDense {
	// if the token in the marking is equal to or greater than the weight of the inhibitory arc (which defaults ot max float64) then assign true otherwise false
	inhibs := p.InhibitoryInputIncidence.(*mat.Dense)
	numPlaces, numTransitions := p.InhibitoryInputIncidence.Dims()
	result := make([]float64, numTransitions)	
	m_0 := mat.VecDenseCopyOf(m.Places)
	
	for i := 0; i < numTransitions; i++ {
		var m_t mat.VecDense
		m_t.SubVec(m_0, inhibs.ColView(i))
		isInhibited := 0.0
		for j := 0; j < numPlaces; j++ {
			if m_t.AtVec(j) >= 0 {
				isInhibited = 1.0
			}
		}
		result[i] = isInhibited
	}

	return mat.NewVecDense(numTransitions, result)
}

func (net *PetriNet) GetEligibleFiringList(marking Marking) (*mat.VecDense, error) {
	numPlaces, numTransitions := net.InputIncidence.Dims()
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
	inhibitedTransitions := ComputeVectorOfInhibitedTransitions(marking, net)
	firingListVector := mat.NewVecDense(numTransitions, firingList)
	var result mat.VecDense
	result.SubVec(firingListVector, inhibitedTransitions)
	return &result, nil
}
func CountEnabledTransitions(firingList *mat.VecDense) int {
	count := 0
	for i := 0; i < firingList.Len(); i++ {
		count += int(firingList.AtVec(i))
	}
	return count
}

func (net *PetriNet) ChooseTransitionFromEligibleFiringList(fullFiringList *mat.VecDense) (*mat.VecDense, error) {
	numTransitions := len(net.TransitionNames)
	firingList := make([]float64, numTransitions, numTransitions)

	// before returning the firing list we need to pick one of the enabled transitions as the only one to be fired
	// first if only one transition is enabled, we return the firing list as is
	// if more than one transition is enabled, we return the firing list with a randomly chosen (free choice) enabled transition set to 1.0
	// and all the other enabled transitions set to 0.0.
	// if no transition is enabled, we have deadlock, but this function can do nothing about it.

	enabledCount := CountEnabledTransitions(fullFiringList)
	if enabledCount == 1 {
		return mat.NewVecDense(numTransitions, firingList), nil
	} else if enabledCount > 1 {
		// randomly choose one of the enabled transitions
		enabledTransitionsToDiscard := rand.Intn(enabledCount + 1)

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

func (net *PetriNet) Fire(m_0 *Marking, firingList *mat.VecDense) (*Marking, error) {
	m_1, err := net.CalculateNextMarking(m_0, firingList)
	if err != nil {
		return m_1, err
	}
	var t TransitionId
	t, err = net.SelectTidOfChosenTransition(firingList)
	if err != nil {
		return m_1, err
	}
	handler := net.TransitionHandlers[t]
	handler(t, m_0, m_1)
	return m_1, nil
}

func (net *PetriNet) CalculateNextMarking(m_0 *Marking, firingList *mat.VecDense) (*Marking, error) {

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

func (net *PetriNet) SelectTidOfChosenTransition(firingList *mat.VecDense) (TransitionId, error) {
	for i := 0; i < firingList.Len(); i++ {
		if firingList.AtVec(i) == 1.0 {
			return TransitionId(i), nil
		}
	}
	return -1, errors.New("net is dead")
}

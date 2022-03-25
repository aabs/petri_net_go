package core

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

/*
 * Building Petri Nets
 */

// For adjacency matrices the layout is places are rows, and transitions are columns
// Terminology:
//	 In, Into:  This means arcs going into a place
//	 Out:  This means arcs going out of a place
type PetriNetBuilder struct {
	Name                     string
	Version                  string
	InputIncidence           []Arc
	OutputIncidence          []Arc
	InhibitoryInputIncidence []Arc
	Places                   []Place
	Transitions              []Transition
	TransitionHandlers       map[TransitionId]func(TransitionId, *Marking, *Marking)
}

func CreatePetriNet() *PetriNetBuilder {
	return &PetriNetBuilder{
		Places:             []Place{},
		Transitions:        []Transition{},
		TransitionHandlers: make(map[TransitionId]func(TransitionId, *Marking, *Marking)),
	}
}

func (p *PetriNetBuilder) Called(name string) *PetriNetBuilder {
	p.Name = name
	return p
}

func (p *PetriNetBuilder) WithVersion(version string) *PetriNetBuilder {
	p.Version = version
	return p
}

func (p *PetriNetBuilder) WithPlace(pl Place) *PetriNetBuilder {
	p.Places = append(p.Places, pl)
	return p
}
func (p *PetriNetBuilder) WithPlaces(places map[PlaceId]string) *PetriNetBuilder {
	for pid, name := range places {
		p.WithPlace(Place{pid, name})
	}
	return p
}
func (p *PetriNetBuilder) WithTransition(t Transition) *PetriNetBuilder {
	p.Transitions = append(p.Transitions, t)
	return p
}
func (p *PetriNetBuilder) WithTransitions(transitions map[TransitionId]string) *PetriNetBuilder {
	for pid, name := range transitions {
		p.WithTransition(Transition{pid, name})
	}
	return p
}
func (p *PetriNetBuilder) WithArcIntoPlace(pid PlaceId, tid TransitionId) *PetriNetBuilder {
	p.InputIncidence = append(p.InputIncidence, Arc{pid, tid})
	return p
}

func (p *PetriNetBuilder) WithArcsIntoPlaces(arcs map[PlaceId][]TransitionId) *PetriNetBuilder {
	//p.InputIncidence = append(p.InputIncidence, Arc{pid, tid})
	for pid, tids := range arcs {
		for _, tid := range tids {
			p.WithArcIntoPlace(pid, tid)
		}
	}
	return p
}

func (p *PetriNetBuilder) WithArcOutOfPlace(tid TransitionId, pid PlaceId) *PetriNetBuilder {
	p.OutputIncidence = append(p.OutputIncidence, Arc{pid, tid})
	return p
}

func (p *PetriNetBuilder) WithArcsOutOfPlaces(arcs map[PlaceId][]TransitionId) *PetriNetBuilder {
	for pid, tids := range arcs {
		for _, tid := range tids {
			p.WithArcOutOfPlace(tid, pid)
		}
	}
	return p
}

func (p *PetriNetBuilder) WithInhibitorArc(pid PlaceId, tid TransitionId) *PetriNetBuilder {
	p.InhibitoryInputIncidence = append(p.InhibitoryInputIncidence, Arc{pid, tid})
	return p
}

func (p *PetriNetBuilder) WithInhibitorArcs(arcs map[PlaceId][]TransitionId) *PetriNetBuilder {
	for pid, tids := range arcs {
		for _, tid := range tids {
			p.WithInhibitorArc(pid, tid)
		}
	}
	return p
}

func (p *PetriNetBuilder) BuildInputIncidenceMatrix() []float64 {
	number_of_places := len(p.Places)
	number_of_transitions := len(p.Transitions)
	result := make([]float64, number_of_places*number_of_transitions)
	// initialize result with zeros
	for i := range result {
		result[i] = 0
	}

	// set input arcs to 1
	for _, arc := range p.InputIncidence {
		row, col := getArrayOffset(arc)
		offset := row*number_of_transitions + (col)
		result[offset] = result[offset] + 1
	}

	return result
}

func getArrayOffset(arc Arc) (int, int) {
	row := int(arc.Place)
	col := int(arc.Transition)
	return row, col
}

func (p *PetriNetBuilder) BuildOutputIncidenceMatrix() []float64 {
	number_of_places := len(p.Places)
	number_of_transitions := len(p.Transitions)
	result := make([]float64, number_of_places*number_of_transitions)
	// initialize result with zeros
	for i := range result {
		result[i] = 0
	}

	// set input arcs to 1
	for _, arc := range p.OutputIncidence {
		row, col := getArrayOffset(arc)
		offset := row*number_of_transitions + (col)
		result[offset] = result[offset] + 1
	}

	return result
}

func (p *PetriNetBuilder) BuildInhibitoryInputIncidenceMatrix() []float64 {
	// A Note on inhibition arcs.
	// Ideally, non-existent inhibitory links should have a weight of infinity, but we are limited to using Max Float64 as an approximation.
	// therefore the weight of an inhibitory arc will always outweight the marking of the pre-place it comes from.

	number_of_places := len(p.Places)
	number_of_transitions := len(p.Transitions)
	result := make([]float64, number_of_places*number_of_transitions)
	// initialize result with zeros
	for i := range result {
		result[i] = math.MaxFloat64
	}

	// set input arcs to 1
	for _, arc := range p.InhibitoryInputIncidence {
		row, col := getArrayOffset(arc)
		offset := row*number_of_transitions + (col)
		if result[offset] == math.MaxFloat64 {
			result[offset] = 1
		} else {
			result[offset] = result[offset] + 1			
		}
	}

	return result
}

func (p *PetriNetBuilder) Build() (*PetriNet, error) {
	number_of_places := len(p.Places)
	number_of_transitions := len(p.Transitions)

	tr_names := make(map[int]string)
	for i, tr := range p.Transitions {
		tr_names[i] = tr.Name
	}
	pl_names := make(map[int]string)
	for i, pl := range p.Places {
		pl_names[i] = pl.Name
	}

	return &PetriNet{
		Name:                     p.Name,
		Version:                  p.Version,
		InputIncidence:           mat.NewDense(number_of_places, number_of_transitions, p.BuildInputIncidenceMatrix()),
		OutputIncidence:          mat.NewDense(number_of_places, number_of_transitions, p.BuildOutputIncidenceMatrix()),
		InhibitoryInputIncidence: mat.NewDense(number_of_places, number_of_transitions, p.BuildInhibitoryInputIncidenceMatrix()),
		PlaceNames:               pl_names,
		TransitionNames:          tr_names,
		TransitionHandlers:       p.TransitionHandlers,
	}, nil
}

func (p *PetriNetBuilder) WithTransitionHandler(tid TransitionId, handler func(tid TransitionId, m_0 *Marking, m_1 *Marking)) *PetriNetBuilder {
	p.TransitionHandlers[tid] = handler
	return p
}

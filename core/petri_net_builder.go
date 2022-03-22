package core

import "gonum.org/v1/gonum/mat"

/*
 * Building Petri Nets
 */

type PetriNetBuilder struct {
	Name                     string
	Version                  string
	InputIncidence           []Pair[int, int]
	OutputIncidence          []Pair[int, int]
	InhibitoryInputIncidence []Pair[int, int]
	Places                   []Place
	Transitions              []Transition
}

func CreatePetriNet() *PetriNetBuilder {
	return &PetriNetBuilder{
		Places:      []Place{},
		Transitions: []Transition{},
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
func (p *PetriNetBuilder) WithTransition(t Transition) *PetriNetBuilder {
	p.Transitions = append(p.Transitions, t)
	return p
}

func (p *PetriNetBuilder) WithInArc(arc Pair[int, int]) *PetriNetBuilder {
	p.InputIncidence = append(p.InputIncidence, arc)
	return p
}

func (p *PetriNetBuilder) WithOutArc(arc Pair[int, int]) *PetriNetBuilder {
	p.OutputIncidence = append(p.OutputIncidence, arc)
	return p
}
func (p *PetriNetBuilder) WithInhibitorArc(arc Pair[int, int]) *PetriNetBuilder {
	p.InhibitoryInputIncidence = append(p.InhibitoryInputIncidence, arc)
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
		result[(arc.Item1-1)*number_of_places+(arc.Item2-1)] = 1
	}

	return result
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
		result[(arc.Item1-1)*number_of_places+(arc.Item2-1)] = 1
	}

	return result
}

func (p *PetriNetBuilder) BuildInhibitoryInputIncidenceMatrix() []float64 {
	number_of_places := len(p.Places)
	number_of_transitions := len(p.Transitions)
	result := make([]float64, number_of_places*number_of_transitions)
	// initialize result with zeros
	for i := range result {
		result[i] = 0
	}

	// set input arcs to 1
	for _, arc := range p.InhibitoryInputIncidence {
		result[(arc.Item1-1)*number_of_places+(arc.Item2-1)] = 1
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
	}, nil
}

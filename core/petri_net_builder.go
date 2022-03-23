package core

import "gonum.org/v1/gonum/mat"

/*
 * Building Petri Nets
 */

// for adjacency matrices the layout is places are rows, and transitions are columns

type PetriNetBuilder struct {
	Name                     string
	Version                  string
	InputIncidence           []Arc
	OutputIncidence          []Arc
	InhibitoryInputIncidence []Arc
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

func (p *PetriNetBuilder) WithInArc(pid PlaceId, tid TransitionId) *PetriNetBuilder {
	p.InputIncidence = append(p.InputIncidence, Arc{pid, tid})
	return p
}

func (p *PetriNetBuilder) WithInArcs(arcs map[PlaceId][]TransitionId) *PetriNetBuilder {
	//p.InputIncidence = append(p.InputIncidence, Arc{pid, tid})
	for pid, tids := range arcs {
		for _, tid := range tids {
			p.WithInArc(pid, tid)
		}
	}
	return p
}

func (p *PetriNetBuilder) WithOutArc(tid TransitionId, pid PlaceId) *PetriNetBuilder {
	p.OutputIncidence = append(p.OutputIncidence, Arc{pid, tid})
	return p
}

func (p *PetriNetBuilder) WithOutArcs(arcs map[PlaceId][]TransitionId) *PetriNetBuilder {
	for pid, tids := range arcs {
		for _, tid := range tids {
			p.WithOutArc(tid, pid)
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
	number_of_places := len(p.Places)
	number_of_transitions := len(p.Transitions)
	result := make([]float64, number_of_places*number_of_transitions)
	// initialize result with zeros
	for i := range result {
		result[i] = 0
	}

	// set input arcs to 1
	for _, arc := range p.InhibitoryInputIncidence {
		row, col := getArrayOffset(arc)
		offset := row*number_of_transitions + (col)
		result[offset] = result[offset] + 1
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

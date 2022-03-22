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

/*
 * Building Places
 */
var placeIdGenerator int = 0

type Place struct {
	Id   int
	Name string
}

type PlaceBuilder struct {
	Id   int
	Name string
}

func CreatePlace() *PlaceBuilder {
	placeIdGenerator++
	return &PlaceBuilder{
		Id: placeIdGenerator,
	}
}

func (p *PlaceBuilder) Called(name string) *PlaceBuilder {
	p.Name = name
	return p
}

func (p *PlaceBuilder) Build() *Place {
	return &Place{
		Id:   p.Id,
		Name: p.Name,
	}
}

type Pair[T1 any, T2 any] struct {
	Item1 T1
	Item2 T2
}

func NewPair[T1 any, T2 any](item1 T1, item2 T2) *Pair[T1, T2] {
	return &Pair[T1, T2]{
		Item1: item1,
		Item2: item2,
	}
}

/*
 * Building Transitions
 */
var transIdGenerator int = 0

type Transition struct {
	Id   int
	Name string
}

type TransitionBuilder struct {
	Id   int
	Name string
}

func CreateTransition() *TransitionBuilder {
	transIdGenerator++
	return &TransitionBuilder{
		Id: transIdGenerator,
	}
}

func (p *TransitionBuilder) Called(name string) *TransitionBuilder {
	p.Name = name
	return p
}

func (p *TransitionBuilder) Build() *Transition {
	return &Transition{
		Id:   p.Id,
		Name: p.Name,
	}
}

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

	return &PetriNet{
		Name:                     p.Name,
		Version:                  p.Version,
		InputIncidence:           mat.NewDense(number_of_places, number_of_transitions, p.BuildInputIncidenceMatrix()),
		OutputIncidence:          mat.NewDense(number_of_places, number_of_transitions, p.BuildOutputIncidenceMatrix()),
		InhibitoryInputIncidence: mat.NewDense(number_of_places, number_of_transitions, p.BuildInhibitoryInputIncidenceMatrix()),
		PlaceNames:               make(map[int]string),
		TransitionNames:          make(map[int]string),
	}, nil
}

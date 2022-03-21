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

func NewPetriNet(number_of_places, number_of_transitions int) (*PetriNet, error) {
	data := make([]float64, number_of_places*number_of_transitions)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	return &PetriNet{
		Name:                     "",
		Version:                  "",
		InputIncidence:           mat.NewDense(number_of_places, number_of_transitions, data),
		OutputIncidence:          mat.NewDense(number_of_places, number_of_transitions, data),
		InhibitoryInputIncidence: mat.NewDense(number_of_places, number_of_transitions, data),
		PlaceNames:               make(map[int]string),
		TransitionNames:          make(map[int]string),
	}, nil
}

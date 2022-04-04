package core

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

type Marking struct {
	InstanceId        string
	DefinitionName    string
	DefinitionVersion string
	Places            mat.Vector
}

func CreateMarking(size int, markings []int) Marking {
	places := mat.NewVecDense(size, ConvertToFloat64(markings))
	return Marking{
		InstanceId:        "unknown instance",
		DefinitionName:    "Unknown Definition",
		DefinitionVersion: "1.0.0",
		Places:            places,
	}
}

func (m *Marking) SetMarking(placeId int, token float64) (*Marking, error) {
	if placeId < 0 || placeId >= m.Places.Len() {
		return nil, errors.New("placeId out of range")
	}
	m_1 := mat.VecDenseCopyOf(m.Places)
	m_1.SetVec(placeId, token)
	return &Marking{Places: m_1}, nil
}

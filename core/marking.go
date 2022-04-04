package core

import "gonum.org/v1/gonum/mat"

type Marking struct {
	InstanceId                string
	DefinitionName    string
	DefinitionVersion string
	Places            mat.Vector
}

func CreateMarking(size int, markings []int) Marking {
	places := mat.NewVecDense(size, ConvertToFloat64(markings))
	return Marking{
		InstanceId:                "unknown instance",
		DefinitionName:    "Unknown Definition",
		DefinitionVersion: "1.0.0",
		Places:            places,
	}
}

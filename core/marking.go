package core

import "gonum.org/v1/gonum/mat"

type Marking struct {
	Places mat.Vector
}

func CreateMarking(size int, markings []int) Marking {
	places := mat.NewVecDense(size, ConvertToFloat64(markings))
	return Marking{
		Places: places,
	}
}

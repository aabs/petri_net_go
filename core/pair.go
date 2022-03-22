package core

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

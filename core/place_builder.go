package core

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

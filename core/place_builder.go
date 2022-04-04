package core

type Place struct {
	Id   PlaceId `json:"id"`
	Name string  `json:"name"`
}

type PlaceBuilder struct {
	Id   PlaceId
	Name string
}

func CreatePlace() *PlaceBuilder {
	return &PlaceBuilder{}
}

func (p *PlaceBuilder) Called(name string) *PlaceBuilder {
	p.Name = name
	return p
}

func (p *PlaceBuilder) WithId(id PlaceId) *PlaceBuilder {
	p.Id = id
	return p
}

func (p *PlaceBuilder) Build() *Place {
	return &Place{
		Id:   p.Id,
		Name: p.Name,
	}
}

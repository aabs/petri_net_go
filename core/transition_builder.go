package core

type Transition struct {
	Id   TransitionId `json:"id"`
	Name string       `json:"name"`
}

type TransitionBuilder struct {
	Id   TransitionId
	Name string
}

func CreateTransition() *TransitionBuilder {
	return &TransitionBuilder{}
}

func (p *TransitionBuilder) Called(name string) *TransitionBuilder {
	p.Name = name
	return p
}

func (p *TransitionBuilder) WithId(id TransitionId) *TransitionBuilder {
	p.Id = id
	return p
}

func (p *TransitionBuilder) Build() *Transition {
	return &Transition{
		Id:   p.Id,
		Name: p.Name,
	}
}

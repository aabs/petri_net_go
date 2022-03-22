package core

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
 
package builder_pattern

// director class manages the multiple builders
// It takes the builder using which client wants
// to generate the House

// todo: should director also implement the builder interface which
// allows the client to build House of his own choice ?
// solve: client can directly make that house using builder.
type director struct{}

func NewDirector() *director {
	return &director{}
}

func (director *director) buildHouse1(concreteBuilder1 *ConcreteBuilder1) *House {
	concreteBuilder1.setDoors(1)
	concreteBuilder1.setFloors(1)
	concreteBuilder1.setRoof("Teekhi")
	return concreteBuilder1.buildHouse()
}

func (director *director) buildHouse2(concreteBuilder2 *ConcreteBuilder2) *House {
	concreteBuilder2.setDoors(2)
	concreteBuilder2.setFloors(2)
	concreteBuilder2.setRoof("Sidhi")
	return concreteBuilder2.buildHouse()
}

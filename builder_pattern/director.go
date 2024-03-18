package builder_pattern

// director class manages the multiple builders
// It takes the builder using which client wants
// to generate the House

// todo: should director also implement the builder interface which
// allows the client to build House of his own choice ?
// solve: client can directly make that house using builder.
type director struct {
	houseBuilder houseBuilder
}

func NewDirector(houseBuilder houseBuilder) *director {
	return &director{
		houseBuilder: houseBuilder,
	}
}

func (director *director) GetHouse() *House {
	director.houseBuilder.setFloors()
	director.houseBuilder.setRoof()
	director.houseBuilder.setDoors()
	return director.houseBuilder.buildHouse()
}

func (director *director) updateBuilder(builder houseBuilder) {
	director.houseBuilder = builder
}

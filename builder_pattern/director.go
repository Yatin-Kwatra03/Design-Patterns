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

func (director *director) buildVilla(villaBuilder *VillaBuilder) *House {
	villaBuilder.setDoors(1)
	villaBuilder.setFloors(1)
	villaBuilder.setRoof("Teekhi")
	return villaBuilder.buildHouse()
}

func (director *director) buildPentHouse(pentHouseBuilder *PentHouseBuilder) *House {
	pentHouseBuilder.setDoors(2)
	pentHouseBuilder.setFloors(2)
	pentHouseBuilder.setRoof("Sidhi")
	return pentHouseBuilder.buildHouse()
}

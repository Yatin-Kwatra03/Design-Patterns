package builder_pattern

// PentHouseBuilder : concrete implementation of a builder
// So let's say PentHouseBuilder creates House of type pent house which
// has some known properties.
type PentHouseBuilder struct {
	floors   int
	roofType string
	doors    int
}

func NewPentHouseBuilder() *PentHouseBuilder {
	return &PentHouseBuilder{}
}

func (builder *PentHouseBuilder) setDoors(doors int) {
	builder.doors = doors
}

func (builder *PentHouseBuilder) setRoof(roofType string) {
	builder.roofType = roofType
}

func (builder *PentHouseBuilder) setFloors(floors int) {
	builder.floors = floors
}

func (builder *PentHouseBuilder) getDoors() int {
	return builder.doors
}

func (builder *PentHouseBuilder) getRoof() string {
	return builder.roofType
}

func (builder *PentHouseBuilder) getFloors() int {
	return builder.floors
}

func (builder *PentHouseBuilder) buildHouse() *House {
	return NewHouse(builder)
}

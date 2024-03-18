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

// compile time safety check
var _ houseBuilder = &PentHouseBuilder{}

func (builder *PentHouseBuilder) setDoors() {
	builder.doors = 2
}

func (builder *PentHouseBuilder) setRoof() {
	builder.roofType = "sidhi"
}

func (builder *PentHouseBuilder) setFloors() {
	builder.floors = 3
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

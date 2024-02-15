package builder_pattern

// ConcreteBuilder2 : concrete implementation of a builder
// So let's say ConcreteBuilder2 creates House of type House2 which
// has some known properties.
type ConcreteBuilder2 struct {
	floors   int
	roofType string
	doors    int
}

func NewConcreteBuilder2() *ConcreteBuilder2 {
	return &ConcreteBuilder2{}
}

func (builder *ConcreteBuilder2) setDoors(doors int) {
	builder.doors = doors
}

func (builder *ConcreteBuilder2) setRoof(roofType string) {
	builder.roofType = roofType
}

func (builder *ConcreteBuilder2) setFloors(floors int) {
	builder.floors = floors
}

func (builder *ConcreteBuilder2) getDoors() int {
	return builder.doors
}

func (builder *ConcreteBuilder2) getRoof() string {
	return builder.roofType
}

func (builder *ConcreteBuilder2) getFloors() int {
	return builder.floors
}

func (builder *ConcreteBuilder2) buildHouse() *House {
	return NewHouse(builder)
}

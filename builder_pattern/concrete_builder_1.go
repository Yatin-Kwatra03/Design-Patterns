package builder_pattern

// ConcreteBuilder1 :
//   - concrete implementation of a builder
//   - this is a popular implementation like
//     villa. So client asks director to build
//     a villa for me.
//
// So let's say ConcreteBuilder1 creates House of type House1 which
// has some known properties.
type ConcreteBuilder1 struct {
	floors   int
	roofType string
	doors    int
}

func NewConcreteBuilder1() *ConcreteBuilder1 {
	return &ConcreteBuilder1{}
}

func (builder *ConcreteBuilder1) setDoors(doors int) {
	builder.doors = doors
}

func (builder *ConcreteBuilder1) setRoof(roofType string) {
	builder.roofType = roofType
}

func (builder *ConcreteBuilder1) setFloors(floors int) {
	builder.floors = floors
}

func (builder *ConcreteBuilder1) getDoors() int {
	return builder.doors
}

func (builder *ConcreteBuilder1) getRoof() string {
	return builder.roofType
}

func (builder *ConcreteBuilder1) getFloors() int {
	return builder.floors
}

func (builder *ConcreteBuilder1) buildHouse() *House {
	return NewHouse(builder)
}

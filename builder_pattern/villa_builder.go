package builder_pattern

// VillaBuilder :
//   - concrete implementation of a villa builder
//   - this is a popular implementation like
//     villa. So client asks director to build
//     a villa for me.
//
// So let's say VillaBuilder creates House of type villa which
// has some known properties.
type VillaBuilder struct {
	floors   int
	roofType string
	doors    int
}

func NewVillaBuilder() *VillaBuilder {
	return &VillaBuilder{}
}

func (builder *VillaBuilder) setDoors(doors int) {
	builder.doors = doors
}

func (builder *VillaBuilder) setRoof(roofType string) {
	builder.roofType = roofType
}

func (builder *VillaBuilder) setFloors(floors int) {
	builder.floors = floors
}

func (builder *VillaBuilder) getDoors() int {
	return builder.doors
}

func (builder *VillaBuilder) getRoof() string {
	return builder.roofType
}

func (builder *VillaBuilder) getFloors() int {
	return builder.floors
}

func (builder *VillaBuilder) buildHouse() *House {
	return NewHouse(builder)
}

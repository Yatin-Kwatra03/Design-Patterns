package builder_pattern

// House : House is the actual entity that client
// needs. Client will need to have access over the
// components of the house like floors, roofType,
// doors etc.
// So in order to do that
// we can also give House interface to client
// and make some fields hidden if we want to
// like architectural design logic. And only
// expose required fields which are of in interest
// to client
type House struct {
	floors   int
	roofType string
	doors    int
}

// NewHouse : creates an actual House entity using
// builder parameters since builder properties
// are set by the client itself.
func NewHouse(houseBuilder houseBuilder) *House {
	return &House{
		floors:   houseBuilder.getFloors(),
		roofType: houseBuilder.getRoof(),
		doors:    houseBuilder.getDoors(),
	}
}

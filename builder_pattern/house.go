package builder_pattern

// House : we can also give House interface to client
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
func NewHouse(builder builder) *House {
	return &House{
		floors:   builder.getFloors(),
		roofType: builder.getRoof(),
		doors:    builder.getDoors(),
	}
}

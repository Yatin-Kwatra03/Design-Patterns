package builder_pattern

// houseBuilder contains all the methods using which
// client can access the components of the house.
type houseBuilder interface {
	setDoors(doors int)
	setRoof(roofType string)
	setFloors(floors int)
	getDoors() int
	getRoof() string
	getFloors() int
	buildHouse() *House
}

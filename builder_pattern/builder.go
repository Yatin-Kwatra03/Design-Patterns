package builder_pattern

type builder interface {
	setDoors(doors int)
	setRoof(roofType string)
	setFloors(floors int)
	getDoors() int
	getRoof() string
	getFloors() int
	buildHouse() *House
}

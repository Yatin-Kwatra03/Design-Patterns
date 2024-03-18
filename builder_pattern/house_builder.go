package builder_pattern

// houseBuilder contains all the methods using which
// client can access the components of the house.
type houseBuilder interface {
	setDoors()
	setRoof()
	setFloors()
	getDoors() int
	getRoof() string
	getFloors() int
	buildHouse() *House
}

func getRelevantHouseBuilder(kiskeJaisa string) houseBuilder {
	if kiskeJaisa == "yatin jaisa" {
		return NewVillaBuilder()
	}

	return NewPentHouseBuilder()
}

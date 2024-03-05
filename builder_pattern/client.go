package builder_pattern

func GharBanwaloKendar(houseType string) *House {
	director := NewDirector()

	if houseType == "yatin jaisa" {
		villaBuilder := NewVillaBuilder()
		return director.buildVilla(villaBuilder)
	}
	pentHouseBuilder := NewPentHouseBuilder()
	return director.buildPentHouse(pentHouseBuilder)
}

// MrjiKaGhar : doesn't uses Director concept at all
func MrjiKaGhar() *House {
	concreteBuilder1 := NewVillaBuilder()
	concreteBuilder1.setFloors(10)
	concreteBuilder1.setRoof("boht uchi")
	concreteBuilder1.setDoors(10)
	return concreteBuilder1.buildHouse()
}

package builder_pattern

func GharBanwaloKendar(houseType string) *House {
	director := NewDirector()

	if houseType == "yatin jaisa" {
		concreteBuilder1 := NewConcreteBuilder1()
		return director.buildHouse1(concreteBuilder1)
	}
	concreteBuilder2 := NewConcreteBuilder2()
	return director.buildHouse2(concreteBuilder2)
}

// MrjiKaGhar : doesn't uses Director concept at all
func MrjiKaGhar() *House {
	concreteBuilder1 := NewConcreteBuilder1()
	concreteBuilder1.setFloors(10)
	concreteBuilder1.setRoof("boht uchi")
	concreteBuilder1.setDoors(10)
	return concreteBuilder1.buildHouse()
}

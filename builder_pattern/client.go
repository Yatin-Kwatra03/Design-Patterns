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
	villaBuilder := NewVillaBuilder()
	villaBuilder.setFloors(10)
	villaBuilder.setRoof("boht uchi")
	villaBuilder.setDoors(10)
	return villaBuilder.buildHouse()
}

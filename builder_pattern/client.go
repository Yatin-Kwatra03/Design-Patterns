package builder_pattern

func GharBanwaloKendar(houseType string) *House {

	houseBuilder := getRelevantHouseBuilder(houseType)
	director := NewDirector(houseBuilder)
	return director.GetHouse()
}

// MrjiKaGhar : doesn't use Director concept at all
func MrjiKaGhar() *House {
	villaBuilder := NewVillaBuilder()
	villaBuilder.setFloors()
	villaBuilder.setRoof()
	villaBuilder.setDoors()
	return villaBuilder.buildHouse()
}

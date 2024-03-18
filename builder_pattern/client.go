package builder_pattern

import "fmt"

func GharBanwaloKendar() {

	houseBuilder1 := getRelevantHouseBuilder("yatin jaisa")
	director := NewDirector(houseBuilder1)
	fmt.Println(director.GetHouse())
	houseBuilder2 := getRelevantHouseBuilder("harleen jaisa")
	director.updateBuilder(houseBuilder2)
	fmt.Println(director.GetHouse())
}

// MrjiKaGhar : doesn't use Director concept at all
func MrjiKaGhar() *House {
	villaBuilder := NewVillaBuilder()
	villaBuilder.setFloors()
	villaBuilder.setRoof()
	villaBuilder.setDoors()
	return villaBuilder.buildHouse()
}

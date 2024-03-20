package prototype

import (
	"fmt"
	"math/rand"
)

func GeneratePrototypeForSoftwareEngineer() {
	// contains all the created clones
	prototypeRegistry := make(map[int]softwareEngineer)

	noobEngineer := newNoobSoftwareEngineer(rand.Int(), "luka doncic", 2)
	prototypeRegistry[noobEngineer.id] = noobEngineer

	experiencedEngineer := newExperiencedSoftwareEngineer(rand.Int(), "kobe bryant", 7)
	prototypeRegistry[experiencedEngineer.id] = experiencedEngineer

	// create clone for a noob software engineer
	noobEngineerPrototypeInterface := prototypeRegistry[noobEngineer.id]
	noobEngineerPrototype, ok_ := noobEngineerPrototypeInterface.(*noobSoftwareEngineer)
	if ok_ {
		fmt.Println("old properties of noob engineer prototype : ", noobEngineerPrototype)
	}
	clonedNoobEngineer := noobEngineerPrototype.Clone().(*noobSoftwareEngineer)
	fmt.Println("cloned noob engineer : ", clonedNoobEngineer)

	// create clone for an experienced software engineer
	// change some values to confirm we create a deep copy
	// and not a shallow copy
	experiencedEngineerPrototypeInterface := prototypeRegistry[experiencedEngineer.id]
	experiencedEngineerPrototype, ok := experiencedEngineerPrototypeInterface.(*experiencedSoftwareEngineer)
	if ok {
		fmt.Println("old properties of experienced engineer prototype : ", experiencedEngineerPrototype)
	}
	clonedExperiencedEngineer := experiencedEngineerPrototype.Clone().(*experiencedSoftwareEngineer)
	clonedExperiencedEngineer.id = 98
	fmt.Println("cloned experienced engineer : ", clonedExperiencedEngineer)

}

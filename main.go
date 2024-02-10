package main

import (
	"fmt"
	"github.com/personal-projects/Design-Patterns/abstract_factory"
	"github.com/personal-projects/Design-Patterns/brute_force"
	"github.com/personal-projects/Design-Patterns/factory"
)

func main() {
	//bruteForceImplementation()
	//factoryMethodImplementation()
	abstractFactoryMethodImplementation()

}

func bruteForceImplementation() {
	brute_force.BruteForceClientExecutionCode()
}

func factoryMethodImplementation() {
	// factory method implementation
	book, err := factory.FactoryClientExecutionCode("sporty")
	if err != nil {
		fmt.Println("error from factory service", err)
		return
	}
	fmt.Println("factory method book info", book.GetBookId(), book.GetBookName())
}

func abstractFactoryMethodImplementation() {
	players := abstract_factory.ParticipateAsPerTheSportsSkill("basketball")
	fmt.Println("need players for basketball : ", players.NoobPlayer.GetNoobPlayer(), ", ", players.ExperiencedPlayer.GetExperiencedPlayer())

	players = abstract_factory.ParticipateAsPerTheSportsSkill("badminton")
	fmt.Println("need players for badminton : ", players.NoobPlayer.GetNoobPlayer(), ", ", players.ExperiencedPlayer.GetExperiencedPlayer())
}

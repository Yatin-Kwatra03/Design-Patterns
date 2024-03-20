package main

import (
	"fmt"

	"github.com/personal-projects/Design-Patterns/abstract_factory"
	"github.com/personal-projects/Design-Patterns/brute_force"
	"github.com/personal-projects/Design-Patterns/builder_pattern"
	"github.com/personal-projects/Design-Patterns/chain_of_responsibility"
	"github.com/personal-projects/Design-Patterns/factory"
	"github.com/personal-projects/Design-Patterns/prototype"
	"github.com/personal-projects/Design-Patterns/singleton"
)

func main() {
	//bruteForceImplementation()
	//factoryMethodImplementation()
	//abstractFactoryMethodImplementation()
	//singletonPatternImplementation()
	//builderPatternImplementation()
	//chainOfResponsibilitiesImplementation()
	prototypeImplementation()
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

func singletonPatternImplementation() {
	// It will be used to check the no of instances created for the singleton class
	singleton.InstanceNo = 0

	singletonInstance1 := singleton.GetCricketTeamInstance()
	fmt.Println(fmt.Sprintf("Instance creation attempt 1 result : %s , instanceNo : %v", singletonInstance1.Name, singleton.InstanceNo))

	singletonInstance2 := singleton.GetCricketTeamInstance()
	fmt.Println(fmt.Sprintf("Instance creation attempt 2 result : %s , instanceNo : %v", singletonInstance2.Name, singleton.InstanceNo))

	// using the old instance to show that it is instantiated only once
	fmt.Println(fmt.Sprintf("Instance creation attempt 3 result : %s , instanceNo : %v", singletonInstance1.Name, singleton.InstanceNo))
}

func builderPatternImplementation() {
	builder_pattern.GharBanwaloKendar()
	fmt.Println(builder_pattern.MrjiKaGhar())
}

func chainOfResponsibilitiesImplementation() {
	sol, err := chain_of_responsibility.ProcessCustomerRequest("login issue")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("complain regarding login issue : ", sol)
	}

	sol, err = chain_of_responsibility.ProcessCustomerRequest("money debited but not credited issue")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("complain regarding money debited but not credited issue : ", sol)
	}

	sol, err = chain_of_responsibility.ProcessCustomerRequest("home screen not loading")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("complain regarding home screen not loading : ", sol)
	}

	sol, err = chain_of_responsibility.ProcessCustomerRequest("balance not updated")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("complain regarding balance not updated : ", sol)
	}

}

func prototypeImplementation() {
	prototype.GeneratePrototypeForSoftwareEngineer()
}

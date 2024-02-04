package main

import (
	"fmt"
	"github.com/personal-projects/Design-Patterns/brute_force"
	"github.com/personal-projects/Design-Patterns/factory"
)

func main() {
	brute_force.BruteForceClientExecutionCode()
	book, err := factory.FactoryClientExecutionCode("sporty")
	if err != nil {
		fmt.Println("error from factory service", err)
		return
	}
	fmt.Println("factory method book info", book.GetBookId(), book.GetBookName())
}

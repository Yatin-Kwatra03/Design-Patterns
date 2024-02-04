package factory

import (
	"errors"
	"fmt"
)

type Factory struct {
	// we are not making these parameters public
	// client won't be able to call it using the factory object
	// client will instead call, what we will expose to them
	// That is the abstraction we want to build over Book creation logic
	biopic  *Biopic
	history *History
}

// FactoryClientExecutionCode : This is the method that will
// be exposed to the client.
// Client just need to share his mood and should get the Book
// for the type

// We are taking example of mood here, but we
// can have some set of parameters on the basis
// of which we can decide the type of book
// we need to generate.

func (s *Factory) GetBookOfType(mood string) (*Book, error) {
	switch mood {
	case "sporty":
		return s.biopic.GetBook()
	case "historic":
		return s.history.GetBook()
	default:
		return nil, errors.New("client has gone crazy, we don't have a book for his mood")
	}
}

func FactoryClientExecutionCode(mood string) {
	// we can consider this as the wire code which
	// handles all the initializations of the services.
	factoryService := initializeFactoryService()

	fmt.Println(fmt.Sprintf("Get me a Book, I am in %s mood today", mood))

	book, err := factoryService.GetBookOfType(mood)
	if err != nil {
		fmt.Println(fmt.Errorf("error while fetching book for %s mood : %w", mood, err))
		return
	}

	fmt.Println(fmt.Sprintf("found the book for the %s mood", mood), book)

}

func initializeFactoryService() *Factory {
	biopicService := NewBiopic()
	historyService := NewHistory()
	return &Factory{
		biopic:  biopicService,
		history: historyService,
	}
}

package factory

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
// we need to generate and we'll return the relevant
// interface. The return type will be of BookInterface
// since we know anyone who implements the interface
// can be initialized using the interface name.

func (s *Factory) getRelevantImpln(mood string) Book {
	if mood == "sporty" {
		return s.biopic
	}
	return s.history
}

func FactoryClientExecutionCode(mood string) (Book, error) {
	// we can consider this as the wire code which
	// handles all the initializations of the services.
	factoryService := initializeFactoryService()
	return factoryService.getRelevantImpln(mood), nil
}

func initializeFactoryService() *Factory {
	biopicService := NewBiopic("yatin di biopic")
	historyService := NewHistory("yatin da itihaas")
	return &Factory{
		biopic:  biopicService,
		history: historyService,
	}
}

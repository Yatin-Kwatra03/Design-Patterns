package chain_of_responsibility

// ProcessCustomerRequest : externally exposed
// method for top layer to trigger complain
// processing flows based on some paramters
func ProcessCustomerRequest(complainType string) (string, error) {
	handlers := make([]baseHandler, 0)
	handlers = append(handlers, newChatbotSupport())
	handlers = append(handlers, newCxExecutive())
	handlers = append(handlers, newTechSupport())

	for idx := 0; idx < len(handlers)-1; idx++ {
		handlers[idx].setNext(handlers[idx+1])
	}

	return handlers[0].handle(&customerRequest{complainType: complainType})
}

package chain_of_responsibility

// ProcessCustomerRequest : externally exposed
// method for top layer to trigger complain
// processing flows based on some paramters
func ProcessCustomerRequest(complainType string) string {
	handler1 := newChatbotSupport()
	handler2 := newCxExecutive()
	handler3 := newTechSupport()
	handler1.setNext(handler2)
	handler2.setNext(handler3)

	return handler1.handle(&customerRequest{complainType: complainType})
}

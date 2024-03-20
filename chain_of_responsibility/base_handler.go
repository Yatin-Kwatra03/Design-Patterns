package chain_of_responsibility

// baseHandler : each handler needs
// to implement the handling logic of
// the request and store reference
// to the next handler (if required)
type baseHandler interface {
	handle(req *customerRequest) (string, error)
	setNext(handler baseHandler)
}

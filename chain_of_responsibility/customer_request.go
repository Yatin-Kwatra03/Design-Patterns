package chain_of_responsibility

// customerRequest : request of the
// customer that needs to be served
// by different levels / handlers.
type customerRequest struct {
	complainType string
}

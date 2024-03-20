package chain_of_responsibility

import "errors"

type cxExecutive struct {
	nextHandler baseHandler
}

func newCxExecutive() *cxExecutive {
	return &cxExecutive{}
}

// compile time safety check
var _ baseHandler = &cxExecutive{}

func (s *cxExecutive) setNext(nextHandler baseHandler) {
	s.nextHandler = nextHandler
}

func (s *cxExecutive) handle(req *customerRequest) (string, error) {
	if req.complainType == "money debited but not credited issue" {
		return "raise dispute", nil
	}
	if s.nextHandler != nil {
		return s.nextHandler.handle(req)
	}
	return "", errors.New("unable to process the request at this time")

}

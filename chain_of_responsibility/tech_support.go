package chain_of_responsibility

import "errors"

type techSupport struct {
	nextHandler baseHandler
}

func newTechSupport() *techSupport {
	return &techSupport{}
}

// compile time safety check
var _ baseHandler = &techSupport{}

func (s *techSupport) setNext(nextHandler baseHandler) {
	s.nextHandler = nextHandler
}

func (s *techSupport) handle(req *customerRequest) (string, error) {
	if req.complainType == "home screen not loading" {
		return "clear cache and reopen the app", nil
	}
	if s.nextHandler != nil {
		return s.nextHandler.handle(req)
	}
	return "", errors.New("unable to process the request at this time")
}

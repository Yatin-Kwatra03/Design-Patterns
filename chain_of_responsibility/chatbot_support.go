package chain_of_responsibility

import "errors"

type chatbotSupport struct {
	nextHandler baseHandler
}

func newChatbotSupport() *chatbotSupport {
	return &chatbotSupport{}
}

// compile time safety check
var _ baseHandler = &chatbotSupport{}

func (s *chatbotSupport) setNext(nextHandler baseHandler) {
	s.nextHandler = nextHandler
}

func (s *chatbotSupport) handle(req *customerRequest) (string, error) {
	if req.complainType == "login issue" {
		return "reinstall the app", nil
	}
	if s.nextHandler != nil {
		return s.nextHandler.handle(req)
	}
	return "", errors.New("unable to process the request at this time")
}

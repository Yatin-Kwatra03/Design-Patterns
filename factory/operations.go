package factory

type Operations interface {
	// GetBook : each book category service will implement the book creation logic as
	// per their need. But that logic should be abstracted from the client.
	GetBook() (*Book, error)
}

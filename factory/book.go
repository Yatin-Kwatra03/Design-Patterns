package factory

type Book interface {
	// GetBookName : each book category service will implement the book creation logic as
	// per their need. But that logic should be abstracted from the client.
	GetBookName() string
	GetBookId() string
}

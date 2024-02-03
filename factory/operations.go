package factory

type StoreOperations interface {
	// GetBooksOfType :  used to fetch the list of books for give a book type
	GetBooksOfType(bookType BookType) ([]*Book, error)

	// GetBookForId : used to get the book corresponding to the given book id
	GetBookForId(id string) (*Book, error)

	// GetBookCountInStore : used to fetch the book count in the store
	GetBookCountInStore() (int32, error)

	// AddBookToStore : used to add new book in the store
	AddBookToStore(book *Book) (*Book, error)
}

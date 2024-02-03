package factory

type Book struct {

	// unique id that will be used to define a book uniquely
	Id string

	// book name
	Name string

	// type of book
	BookType BookType

	// price of the book
	Price int32
}

type BookType int32

const (
	BookType_BOOK_TYPE_UNSPECIFIED BookType = 0
	BookType_BOOK_TYPE_HISTORY     BookType = 1
	BookType_BOOK_TYPE_PHILOSOPHY  BookType = 2
	BookType_BOOK_TYPE_BIOPIC      BookType = 0
	BookType_BOOK_TYPE_MYTHOLOGY   BookType = 0
)

package factory

// Book : Different services will implement these functions, which will be used by clients
// to fetch the book info.
type Book interface {
	GetBookName() string
	GetBookId() string
}

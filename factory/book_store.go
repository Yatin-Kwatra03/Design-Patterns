package factory

import (
	"errors"
	"math/rand"
	"time"
)

type BookStore struct {

	// StoreId : unique id for the book store
	StoreId string

	// Name : name of the book store
	Name string

	// Location : location  / Address of the book store
	Location string

	// PhoneNumber : phone number of the book store
	PhoneNumber string

	// WebUrlForShop : online web url of the book store
	WebUrlForShop string

	// BookShelf : stores all the book for a given book type

	// Based on the requirement , we can build the database schema and choose Indexes.
	// Here major requirement was to fetch book for a given book type, so added a map
	// from book type to book .
	BookShelf map[BookType][]*Book
}

func NewBookStore(name, location, phoneNumber, webUrlForShop string) (*BookStore, error) {
	bookShelf := make(map[BookType][]*Book)
	rand.Seed(time.Now().UnixNano())

	return &BookStore{
		StoreId:       string(rand.Intn(10000000)),
		Name:          name,
		Location:      location,
		PhoneNumber:   phoneNumber,
		WebUrlForShop: webUrlForShop,
		BookShelf:     bookShelf,
	}, nil
}

func (s *BookStore) AddBookToStore(book *Book) (*Book, error) {
	if book == nil || book.BookType == BookType_BOOK_TYPE_UNSPECIFIED {
		return nil, errors.New("invalid book object passed in request")
	}

	rand.Seed(time.Now().UnixNano())
	book.Id = string(rand.Intn(10000000))

	// add book to the book shelf
	s.BookShelf[book.BookType] = append(s.BookShelf[book.BookType], book)
	return book, nil
}

func (s *BookStore) GetBooksOfType(bookType BookType) ([]*Book, error) {
	if bookType == BookType_BOOK_TYPE_UNSPECIFIED {
		return nil, errors.New("book type cannot be unspecified")
	}

	booksOfGivenType, ok := s.BookShelf[bookType]
	if !ok {
		return nil, errors.New("no book found of given type")
	}
	return booksOfGivenType, nil
}

func (s *BookStore) GetBookForId(id string) (*Book, error) {
	return nil, nil
}

func (s *BookStore) GetBookCountInStore() (int32, error) {
	return 0, nil
}

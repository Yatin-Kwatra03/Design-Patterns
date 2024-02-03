package main

import (
	"fmt"
	"github.com/epifi/gamma/pkg/logger"
	"github.com/personal-projects/Design-Patterns/factory"
)

func main() {
	fmt.Println("yeah")
	bookStore, err := factory.NewBookStore("book-mehakma", "kaggadaspure", "919988221122", "gareebDiBooksLelo.com")
	if err != nil {
		logger.ErrorNoCtx("error while initialising a new book store")
		return
	}

	createdBookEntity, err := bookStore.AddBookToStore(&factory.Book{
		Name:     "jordan-k-raaz",
		BookType: factory.BookType_BOOK_TYPE_BIOPIC,
		Price:    150,
	})
	if err != nil {
		logger.ErrorNoCtx("error in creating book entity")
		return
	}

	createdBookEntity, err = bookStore.AddBookToStore(&factory.Book{
		Name:     "kobe-k-raaz",
		BookType: factory.BookType_BOOK_TYPE_BIOPIC,
		Price:    250,
	})
	if err != nil {
		logger.ErrorNoCtx("error in creating book entity")
		return
	}

	fmt.Println("book added to store ", createdBookEntity)

	fetchedBooks, err := bookStore.GetBooksOfType(factory.BookType_BOOK_TYPE_BIOPIC)
	if err != nil {
		logger.ErrorNoCtx("error in fetching books of type biopic")
		return
	}
	fmt.Println("fetched biopics from book store", fetchedBooks)
}

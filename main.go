package main

import (
	"fmt"
	"github.com/personal-projects/Design-Patterns/factory"
)

func main() {
	bookStore, err := factory.NewBookStore("book-mehakma", "kaggadaspure", "919988221122", "gareebDiBooksLelo.com")
	if err != nil {
		fmt.Println("error while initialising a new book store : ", err)
		return
	}

	createdBookEntity, err := bookStore.AddBookToStore(&factory.Book{
		Name:     "jordan-k-raaz",
		BookType: factory.BookType_BOOK_TYPE_BIOPIC,
		Price:    150,
	})
	if err != nil {
		fmt.Println("error in creating book entity", createdBookEntity.Name)
		return
	}

	createdBookEntity, err = bookStore.AddBookToStore(&factory.Book{
		Name:     "kobe-k-raaz",
		BookType: factory.BookType_BOOK_TYPE_BIOPIC,
		Price:    250,
	})
	if err != nil {
		fmt.Println("error in creating book entity", createdBookEntity.Name)
		return
	}

	fmt.Println("book added to store ", createdBookEntity)

	fetchedBooks, err := bookStore.GetBooksOfType(factory.BookType_BOOK_TYPE_BIOPIC)
	if err != nil {
		fmt.Println("error in fetching books of type biopic", err)
		return
	}
	fmt.Println("fetched biopics from book store", fetchedBooks[0].Name, fetchedBooks[1].Name)
}

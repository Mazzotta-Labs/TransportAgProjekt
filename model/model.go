package model

import (
	"os"
)

/*
func AddBook(book entity.Book) {
	internal.R.AddBook(book)
}

func CheckBookAvailability(isbn string) bool {

	for _, book := range internal.R.FindAll() {
		if book.ISBN == isbn && !book.Borrowed {
			return true
		}
	}
	return false
}

func LendBook(isbn string) *entity.Book {
	for _, book := range internal.R.FindAll() {
		if book.ISBN == isbn && !book.Borrowed {
			book.Borrowed = true
			internal.R.UpdateBook(book)
			return &book
		}
	}
	return nil
}

func ReturnBook(isbn string) *entity.Book {
	for _, book := range internal.R.FindAll() {
		if book.ISBN == isbn && book.Borrowed {
			book.Borrowed = false
			internal.R.UpdateBook(book)
			return &book
		}
	}
	return nil
}

func FindAllBooks() []entity.Book {
	return internal.R.FindAll()
}
*/

// Model
func ShutDown() {
	os.Exit(0)
}

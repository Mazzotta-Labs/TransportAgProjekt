package model

import (
	"TransportAgProjekt1/model/internal"
	"mvc-booklibrary/book/entity"
	"os"
)

func AddBook(book entity.Book) {
	/*TODO
	internal.R.AddBook(book)

	 */
}

func CheckBookAvailability(isbn string) bool {
	/*TODO
	for _, book := range internal.R.FindAll() {
		if book.ISBN == isbn && !book.Borrowed {
			return true
		}
	}
	return false

	 */
}

func LendBook(isbn string) *entity.Book {
	/*TODO
	for _, book := range internal.R.FindAll() {
		if book.ISBN == isbn && !book.Borrowed {
			book.Borrowed = true
			internal.R.UpdateBook(book)
			return &book
		}
	}
	return nil

	 */
}

func ReturnBook(isbn string) *entity.Book {
	/*TODO
	for _, book := range internal.R.FindAll() {
		if book.ISBN == isbn && book.Borrowed {
			book.Borrowed = false
			internal.R.UpdateBook(book)
			return &book
		}
	}
	return nil
	*/
}

func FindAllBooks() []entity.Book {
	/*TODO
	return internal.R.FindAll()
	 */
}

// Model
func ShutDown() {
	os.Exit(0)
}

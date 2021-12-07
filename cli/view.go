package cli

import (
	"fmt"
	"mvc-booklibrary/book/entity"
)

// view

func PrintMenue() {
	fmt.Println(`
###########################################
#******* WELCOME TO OUR LIBRARY ***********
#******* CHOOSE YOUR OPTION BELOW *********
# 1. ADD ANOTHER BOOK IN LIBRARY
# 2. REMOVE A BOOK FROM A LIBRARY
# 3. CHECK AVAILABILITY
# 4. LEND A BOOK
# 5. RETURN A BOOK
# 6. VIEW ALL BOOKS
#
# q. TERMINATE BOOK LIBRARY APP
`)
}

func printBookInformation() {
	fmt.Println(`
Please enter all information based on this pattern:
Pattern: ISBN, Title, Author, Publishing Year:
`)
}

func printEnterIsbnNumber() {
	fmt.Println(`
Please enter the ISBN number:
`)
}

func printIsBookAvailable(isAvailable bool) {
	if isAvailable {
		fmt.Println("Yes, the Book is available!")
	} else {
		fmt.Println("No, the Book is not available!")
	}
	printContinue()
}

func printBookList(booksToPrint []entity.Book) {
	for i, book := range booksToPrint {
		fmt.Println(i+1, "| ISBN:", book.ISBN+",", "TITLE:", book.Title+",", "AUTHOR:", book.Author+",", "BORROWED:", book.Borrowed)
	}
	printContinue()
}


func printIsBookBorrowed(book *entity.Book) {
	if book == nil {
		fmt.Println("Sorry, model not available!")
		printContinue()
		return
	}

	fmt.Println("Book:", book.ISBN, book.Title, "borrowed")
	printContinue()
}

func printIsBookReturned(book *entity.Book) {
	if book == nil {
		fmt.Println("Sorry, we are not expecting this model to be returned!")
		printContinue()
		return
	}

	fmt.Println("Book:", book.ISBN, book.Title, "returned")
	printContinue()
}

func printContinue() (int, error) {
	return fmt.Println("Press c to continue!")
}

func printGodBye() {
	fmt.Println("Good bye!")
}


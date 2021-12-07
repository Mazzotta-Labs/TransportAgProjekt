package cli

import (
	"bufio"
	"log"
	book2 "mvc-booklibrary/book"
	"mvc-booklibrary/book/entity"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func parseCommand(input string) {
	switch {
	case input == "1":
		ClearTerminal()
		printBookInformation()
		response := askForCommand()
		book := createBook(response)
		book2.AddBook(*book)
		ClearTerminal()
		PrintMenue()
		break
	case input == "3":
		ClearTerminal()
		printEnterIsbnNumber()
		isbn := askForCommand()
		ClearTerminal()
		isAvailable := book2.CheckBookAvailability(isbn)
		printIsBookAvailable(isAvailable)
		break
	case input == "4":
		ClearTerminal()
		printEnterIsbnNumber()
		isbn := askForCommand()
		ClearTerminal()
		book := book2.LendBook(isbn)
		printIsBookBorrowed(book)
		break
	case input == "5":
		ClearTerminal()
		printEnterIsbnNumber()
		isbn := askForCommand()
		ClearTerminal()
		book := book2.ReturnBook(isbn)
		printIsBookReturned(book)
		break
	case input == "6":
		ClearTerminal()
		books := book2.FindAllBooks()
		printBookList(books)
		break
	case input == "c":
		ClearTerminal()
		PrintMenue()
		break
	case input == "q":
		ClearTerminal()
		printGodBye()
		book2.ShutDown()
		break
	}
}

func createBook(response string) *entity.Book {
	bookInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")
	return &entity.Book{
		ISBN:          bookInfos[0],
		Title:         bookInfos[1],
		Author:        bookInfos[2],
		PublishedYear: toInt(bookInfos[3]),
		Borrowed:      false,
	}
}

func askForCommand() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSpace(response)
	return response
}

func ExecuteCommand() {
	command := askForCommand()
	parseCommand(command)
}

func ClearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}


func toInt(info string) int {
	aInt, _ := strconv.Atoi(info)
	return aInt
}
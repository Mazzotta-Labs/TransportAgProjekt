package cli

import (
	"TransportAgProjekt1/model"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func parseCommand(input string) {
	switch {
	case input == "1":
		ClearTerminal()
		PrintOrderMenu()
	out1:
		for true {
			command := askForCommand()
			switch {
			case command == "1":
				//TODO
				break
			case command == "2":
				//TODO
				break
			case command == "3":
				//TODO
				break
			case command == "q":
				break out1
			default:
				fmt.Println("ungültige Eingabe")
			}
		}
		ClearTerminal()
		PrintMenue()
		break
	case input == "2":
		ClearTerminal()
		PrintDriverMenu()
	out2:
		for true {
			command := askForCommand()
			switch {
			case command == "1":
				//TODO
				break
			case command == "2":
				//TODO
				break
			case command == "3":
				//TODO
				break
			case command == "q":
				break out2
			default:
				fmt.Println("ungültige Eingabe")
			}
		}
		ClearTerminal()
		PrintMenue()
		break
	case input == "3":
		ClearTerminal()
		PrintCustomerMenu()
	out3:
		for true {
			command := askForCommand()
			switch {
			case command == "1":
				//TODO
				break
			case command == "2":
				//TODO
				break
			case command == "3":
				//TODO
				break
			case command == "q":
				break out3
			default:
				fmt.Println("ungültige Eingabe")
			}
		}
		ClearTerminal()
		PrintMenue()
		break
	case input == "4":
		ClearTerminal()
		PrintProductMenu()
	out4:
		for true {
			command := askForCommand()
			switch {
			case command == "1":
				//TODO
				break
			case command == "2":
				//TODO
				break
			case command == "3":
				//TODO
				break
			case command == "q":
				break out4
			default:
				fmt.Println("ungültige Eingabe")
			}
		}
		ClearTerminal()
		PrintMenue()
		break
	case input == "q":
		ClearTerminal()
		printGodBye()
		model.ShutDown()
		break
	default:
		println("ungültige Eingabe")
	}
}

/*
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

*/

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

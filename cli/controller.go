package cli

import (
	"TransportAgProjekt/model"
	"TransportAgProjekt/model/entity"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func parseCommand(input string) {
	switch {
	case input == "0":
		test()
	case input == "1":
		ClearTerminal()
		PrintOrderMenu()
		orders := model.FindAllOrder()
		printOrderList(orders)
	out1:
		for true {
			command := askForCommand()
			switch {
			case command == "1":
				PrintAddOrder()
				command := askForCommand()
				order := createOrder(command)
				model.AddOrder(*order)
				break
			case command == "2":
				PrintUpdateOrder()
				command := askForCommand()
				order := createOrder(command)
				model.UpdateOrder(*order)
				break
			case command == "3":
				PrintDeleteOrder()
				command := askForCommand()
				model.DeleteOrder(command)
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
	out2:
		for true {
			ClearTerminal()
			PrintDriverMenu()
			drivers := model.FindAllDriver()
			printDriverList(drivers)
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
	out3:
		for true {
			ClearTerminal()
			PrintCustomerMenu()
			result := model.FindAllCustomer()
			printCustomerList(result)
			command := askForCommand()
			switch {
			case command == "1":
				PrintAddCustomer()
				command := askForCommand()
				customer := createCustomer(command)
				model.AddCustomer(*customer)
				break
			case command == "2":
				PrintUpdateCustomer()
				command := askForCommand()
				customer := createCustomer(command)
				model.UpdateCustomer(*customer)
				break
			case command == "3":
				PrintDeleteCustomer()
				command := askForCommand()
				model.DeleteCustomer(command)
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

func createCustomer(response string) *entity.Customer {
	customerInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")
	return &entity.Customer{
		CustomerId:      toInt(customerInfos[0]),
		CustomerName:    customerInfos[1],
		CustomerPrename: customerInfos[2],
		TelNr:           customerInfos[3],
		Street:          customerInfos[4],
		Plz:             customerInfos[5],
		Town:            customerInfos[6],
		Country:         customerInfos[7],
	}
}

func createOrder(response string) *entity.Order {
	orderInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")
	dateStamp, err := time.Parse("2006-01-02", orderInfos[1])
	intCustomerVar, err := strconv.Atoi(orderInfos[2])
	intDriverVar, err := strconv.Atoi(orderInfos[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &entity.Order{
		OrderId:    toInt(orderInfos[0]),
		OrderDate:  dateStamp.Local(),
		CustomerId: intCustomerVar,
		DriverId:   intDriverVar,
	}
}

func printDriverList(driversToPrint []entity.Driver) {
	for i, driver := range driversToPrint {
		fmt.Println(i+1,
			"| Fahrer ID:", driver.DriverId+",",
			"Name:", driver.Name+",",
			"Vorname:", driver.Prename+",",
			"Fahreug ID:", driver.VehicleId+",",
			"Fahrzeug Marke:", driver.Brand+",",
			"Fahrzeugnummer:", driver.Number)
	}
}

func printCustomerList(customersToPrint []entity.Customer) {
	for i, customer := range customersToPrint {
		fmt.Println(i+1, "| Kunden ID:", toStr(customer.CustomerId)+",",
			"Name:", customer.CustomerName+",",
			"Vorname:", customer.CustomerPrename+",",
			"Tel:", customer.TelNr+",",
			"PLZ:", customer.Plz+",",
			"Ort:", customer.Town,
			"Strasse:", customer.Street,
			"Land:", customer.Country)
	}
}

func printOrderList(ordersToPrint []entity.Order) {
	for i, order := range ordersToPrint {
		fmt.Println(i+1, "| Order ID:", toStr(order.OrderId)+",",
			"Datum:", order.OrderDate.Format("RFC822")+",",
			"Kunden Id:", strconv.Itoa(order.CustomerId)+",",
			"Fahrer Id:", strconv.Itoa(order.DriverId)+",")
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

func toStr(info int) string {
	aStr := strconv.Itoa(info)
	return aStr
}

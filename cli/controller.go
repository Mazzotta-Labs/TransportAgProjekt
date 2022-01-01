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
	case input == "1":
	out1:
		for true {
			ClearTerminal()
			PrintOrderMenu()
			orders := model.FindAllOrder()
			printOrderList(orders)
			command := askForCommand()
			switch {
			case command == "1":
				PrintAddOrder()
				command := askForCommand()
				order := createOrder(command, true)
				println(&order)
				model.AddOrder(*order)
				break
			case command == "2":
				PrintUpdateOrder()
				command := askForCommand()
				order := createOrder(command, false)
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
				PrintAddDriver()
				command := askForCommand()
				driver := createDriver(command, true)
				model.AddDriver(*driver)
				break
			case command == "2":
				PrintUpdateDriver()
				command := askForCommand()
				driver := createDriver(command, false)
				model.UpdateDriver(*driver)
				break
			case command == "3":
				PrintDeleteDriver()
				command := askForCommand()
				model.DeleteDriver(command)
				ClearTerminal()
				PrintDriverMenu()
				drivers := model.FindAllDriver()
				printDriverList(drivers)
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
	out3: //Marke da aus 2 Schleifen in einander ausgestiegen werden muss
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
				customer := createCustomer(command, true)
				model.AddCustomer(*customer)
				break
			case command == "2":
				PrintUpdateCustomer()
				command := askForCommand()
				customer := createCustomer(command, false)
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
				println("ungültige Eingabe")
				time.Sleep(2 * time.Second)
			}
		}
		ClearTerminal()
		PrintMenue()
		break
	case input == "4":
	out4:
		for true {
			ClearTerminal()
			PrintProductMenu()
			products := model.FindAllProduct()
			printProductList(products)
			command := askForCommand()
			switch {
			case command == "1":
				PrintAddProduct()
				command := askForCommand()
				product := createProduct(command, true)
				model.AddProduct(*product)
				break
			case command == "2":
				PrintUpdateProduct()
				command := askForCommand()
				product := createProduct(command, false)
				model.UpdateProduct(*product)
				break
			case command == "3":
				PrintDeleteProduct()
				command := askForCommand()
				model.DeleteProduct(command)
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
	case input == "5":
		ClearTerminal()
		PrintVehicleMenu()
		vehicles := model.FindAllVehicles()
		printVehicleList(vehicles)
	out5:
		for true {
			command := askForCommand()
			switch {
			case command == "1":
				PrintAddVehicle()
				command := askForCommand()
				vehicle := createVehicle(command, true)
				model.AddVehicles(*vehicle)
				break
			case command == "2":
				PrintUpdateVehicle()
				command := askForCommand()
				vehicle := createVehicle(command, false)
				model.UpdateVehicles(*vehicle)
				break
			case command == "3":
				PrintDeleteVehicle()
				command := askForCommand()
				model.DeleteVehicles(command)
				break
			case command == "q":
				break out5
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
		time.Sleep(2 * time.Second)
	}
}

func createCustomer(response string, isCreate bool) *entity.Customer {
	customerInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")
	if isCreate {
		return &entity.Customer{
			CustomerName:    customerInfos[0],
			CustomerPrename: customerInfos[1],
			TelNr:           customerInfos[2],
			Street:          customerInfos[3],
			Plz:             customerInfos[4],
			Town:            customerInfos[5],
			Country:         customerInfos[6],
		}
	}
	return &entity.Customer{
		CustomerId:      0,
		CustomerName:    customerInfos[0],
		CustomerPrename: customerInfos[1],
		TelNr:           customerInfos[2],
		Street:          customerInfos[3],
		Plz:             customerInfos[4],
		Town:            customerInfos[5],
		Country:         customerInfos[6],
	}
}

func createVehicle(response string, isCreate bool) *entity.Vehicle {
	customerInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")
	if isCreate {
		return &entity.Vehicle{
			Brand:  customerInfos[0],
			Number: customerInfos[1],
		}
	}
	return &entity.Vehicle{
		VehicleId: toInt(customerInfos[0]),
		Brand:     customerInfos[1],
		Number:    customerInfos[2],
	}
}

func createDriver(response string, isCreate bool) *entity.Driver {
	driverInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")
	if isCreate {
		return &entity.Driver{
			Name:      driverInfos[0],
			Prename:   driverInfos[1],
			VehicleId: toInt(driverInfos[2]),
		}
	}

	return &entity.Driver{
		DriverId:  toInt(driverInfos[0]),
		Name:      driverInfos[1],
		Prename:   driverInfos[2],
		VehicleId: toInt(driverInfos[3]),
	}
}

func createOrder(response string, isCreate bool) *entity.Order {
	orderInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")
	indexRange := 0
	if isCreate {
		indexRange = 1
	}

	intCustomerVar, err := strconv.Atoi(orderInfos[2-indexRange])
	intDriverVar, err := strconv.Atoi(orderInfos[3-indexRange])
	var intProductsVar []int
	for i := range orderInfos {
		if i > 2 {
			product, err := strconv.Atoi(orderInfos[i])
			intProductsVar = append(intProductsVar, product)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if isCreate {
		return &entity.Order{
			OrderDate:  orderInfos[0],
			CustomerId: intCustomerVar,
			DriverId:   intDriverVar,
			ProductsId: intProductsVar,
		}
	}

	return &entity.Order{
		OrderId:    toInt(orderInfos[0]),
		OrderDate:  orderInfos[1],
		CustomerId: intCustomerVar,
		DriverId:   intDriverVar,
	}
}

func createProduct(response string, isCreate bool) *entity.Product {
	productInfos := strings.Split(strings.ReplaceAll(response, ", ", ","), ",")

	if isCreate {
		return &entity.Product{
			Description: productInfos[0],
			CategoryId:  toInt(productInfos[1]),
			Name:        productInfos[2],
		}
	}

	return &entity.Product{
		ProductId:   toInt(productInfos[0]),
		Description: productInfos[1],
		CategoryId:  toInt(productInfos[2]),
		Name:        productInfos[3],
	}
}

func printDriverList(driversToPrint []entity.Driver) {
	for i, driver := range driversToPrint {
		fmt.Println(i+1,
			"| Fahrer ID:", toStr(driver.DriverId)+",",
			"Name:", driver.Name+",",
			"Vorname:", driver.Prename+",",
			"Fahreug ID:", toStr(driver.VehicleId)+",",
			"Fahrzeug Marke:", driver.Brand+",",
			"Fahrzeugnummer:", driver.Number)
	}
}

func printVehicleList(vehiclesToPrint []entity.Vehicle) {
	for i, vehicle := range vehiclesToPrint {
		fmt.Println(i+1,
			"| Fahrzeug ID:", toStr(vehicle.VehicleId)+",",
			"Marke:", vehicle.Brand+",",
			"Fahrername:", vehicle.DriverId+",",
			"Fahreug Nummer:", vehicle.Number+",")
	}
}

func printCustomerList(customersToPrint []entity.Customer) {
	for i, customer := range customersToPrint {
		fmt.Println(i+1, "| Kunden ID:", toStr(customer.CustomerId)+",",
			"Name:", customer.CustomerName+",",
			"Vorname:", customer.CustomerPrename+",",
			"Tel:", customer.TelNr+",",
			"Strasse:", customer.Street,
			"PLZ:", customer.Plz+",",
			"Ort:", customer.Town,
			"Land:", customer.Country)
	}
}

func printOrderList(ordersToPrint []entity.Order) {
	for i, order := range ordersToPrint {
		date, _ := time.Parse("2006-01-02", order.OrderDate)
		fmt.Println(i+1, "| Order ID:", toStr(order.OrderId)+",",
			"Datum:", date.Format("2006-01-02")+",",
			"Kunden Id:", strconv.Itoa(order.CustomerId)+",",
			"Fahrer Id:", strconv.Itoa(order.DriverId)+",",
			"Produkt Id:", strconv.Itoa(order.ProductId))
	}
}

func printProductList(productsToPrint []entity.Product) {
	for i, product := range productsToPrint {
		fmt.Println(i+1, "| Product ID:", toStr(product.ProductId)+",",
			"Beschreibung:", product.Description+",",
			"Kategorie Id:", strconv.Itoa(product.CategoryId)+",",
			"Name:", product.Name)
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

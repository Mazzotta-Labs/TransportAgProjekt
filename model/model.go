package model

import (
	"TransportAgProjekt/model/entity"
	"TransportAgProjekt/model/internal"
	"os"
)

func FindAllDriver() []entity.Driver {
	return internal.R.FindAllDriver()
}

func AddDriver(driver entity.Driver) {
	internal.R.AddDriver(driver)
}

func UpdateDriver(driver entity.Driver) {
	internal.R.UpdateDriver(driver)
}

func DeleteDriver(driver entity.Driver) {
	internal.R.DeleteDriver(driver)
}

func FindAllCustomer() []entity.Customer {
	return internal.R.FindAllCustomer()
}

func AddCustomer(customer entity.Customer) {
	internal.R.AddCustomer(customer)
}

func UpdateCustomer(customer entity.Customer) {
	internal.R.UpdateCustomer(customer)
}

func DeleteCustomer(customer string) {
	internal.R.DeleteCustomer(customer)
}

func FindAllOrder() []entity.Order {
	return internal.R.FindAllOrder()
}

func AddOrder(order entity.Order) {
	internal.R.AddOrder(order)
}

func UpdateOrder(order entity.Order) {
	internal.R.UpdateOrder(order)
}

func DeleteOrder(orderId string) {
	internal.R.DeleteOrder(orderId)
}

func FindAllProduct() []entity.Product {
	return internal.R.FindAllProduct()
}

func AddProduct(product entity.Product) {
	internal.R.AddProduct(product)
}

func UpdateProduct(product entity.Product) {
	internal.R.UpdateProduct(product)
}

func DeleteProduct(productId string) {
	internal.R.DeleteProduct(productId)
}

// Vehicles

func FindAllVehicles() []entity.Vehicle {
	return internal.R.FindAllVehicle()
}

func AddVehicles(vehicle entity.Vehicle) {
	internal.R.AddVehicle(vehicle)
}

func UpdateVehicles(vehicle entity.Vehicle) {
	internal.R.UpdateVehicle(vehicle)
}

func DeleteVehicles(vehicle entity.Vehicle) {
	internal.R.DeleteVehicle(vehicle)
}

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

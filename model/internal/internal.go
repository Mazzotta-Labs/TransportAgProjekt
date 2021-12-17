package internal

import (
	"TransportAgProjekt/model/entity"
	"TransportAgProjekt/model/internal/mysql"
)

var R = mysql.NewMySqlRepository()

type Repository interface {
	FindAllDriver() []entity.Driver
	AddDriver(driver entity.Driver)
	UpdateDriver(driver entity.Driver)
	DeleteDriver(driver entity.Driver)
	FindAllCustomer() []entity.Customer
	AddCustomer(customer entity.Customer)
	UpdateCustomer(customer entity.Customer)
	DeleteCustomer(customer string)
	FindAllOrder() []entity.Order
	AddOrder(order entity.Order)
	UpdateOrder(order entity.Order)
	DeleteOrder(order entity.Order)
	FindAllProduct() []entity.Product
	AddProduct(product entity.Product)
	UpdateProduct(product entity.Product)
	DeleteProduct(product entity.Product)
}

/*
var R = mysql.NewMySqlRepository()


type Repository interface {
	FindAll() []entity.Book
	AddBook(book entity.Book)
	UpdateBook(book entity.Book)
}

*/

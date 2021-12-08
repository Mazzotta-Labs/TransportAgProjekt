package internal

import (
	"TransportAgProjekt1/model/entity"
	"TransportAgProjekt1/model/internal/mysql"
)

var R = mysql.NewMySqlRepository()

type Repository interface {
	FindAllDriver() []entity.Driver
	AddBookDriver(driver entity.Driver)
	UpdateDriver(driver entity.Driver)
	DeleteDriver(driver entity.Driver)
}

/*
var R = mysql.NewMySqlRepository()


type Repository interface {
	FindAll() []entity.Book
	AddBook(book entity.Book)
	UpdateBook(book entity.Book)
}

*/

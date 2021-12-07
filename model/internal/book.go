package internal

import (
	"mvc-booklibrary/book/entity"
	"mvc-booklibrary/book/internal/mysql"
)

var R = mysql.NewMySqlRepository()


type Repository interface {
	FindAll() []entity.Book
	AddBook(book entity.Book)
	UpdateBook(book entity.Book)
}
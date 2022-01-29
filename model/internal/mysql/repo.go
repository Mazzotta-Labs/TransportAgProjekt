package mysql

import (
	"TransportAgProjekt/model/entity"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type MySqlRepository struct {
	orders []entity.Order
}

func NewMySqlRepository() *MySqlRepository {
	repository := MySqlRepository{}
	db = openDatabase()
	repository.orders = repository.FindAllOrder()
	return &repository
}

func openDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(192.168.1.12:3306)/transportag")
	if err != nil {
		panic(err.Error())
	}
	println("DB connected")
	return db
}

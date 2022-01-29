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
	db, err := sql.Open("mysql", "student-dnd-db:7nS$!!8T@tcp(localhost:3306)/transportag")
	if err != nil {
		panic(err.Error())
	}
	return db
}

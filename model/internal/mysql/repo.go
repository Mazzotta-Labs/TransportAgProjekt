package mysql

import (
	"TransportAgProjekt/model/entity"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
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
	db, err := sql.Open("mysql", "root:7nS$!!8T@tcp(student-dnd-vm:3306)/transportag")
	if err != nil {
		println(err)
		time.Sleep(10 * time.Second)
		//panic(err.Error())
	}
	return db
}

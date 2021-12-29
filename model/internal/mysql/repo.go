package mysql

import (
	"TransportAgProjekt/model/entity"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type MySqlRepository struct {
	drivers   []entity.Driver
	customers []entity.Customer
	orders    []entity.Order
	products  []entity.Product
}

func NewMySqlRepository() *MySqlRepository {
	repository := MySqlRepository{}
	db = openDatabase()
	repository.drivers = repository.FindAllDriver()
	repository.customers = repository.FindAllCustomer()
	repository.orders = repository.FindAllOrder()
	repository.products = repository.FindAllProduct()
	return &repository
}

func openDatabase() *sql.DB {
	fmt.Println("Bitte MySQL ConnectionString eingeben (Format: root:root@tcp(localhost:3306)/transportag)")

	command := askForCommand()

	db, err := sql.Open("mysql", command)
	if err != nil {
		panic(err.Error())
	}
	return db
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

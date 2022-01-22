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
	orders []entity.Order
}

func NewMySqlRepository() *MySqlRepository {
	repository := MySqlRepository{}
	db = openDatabase()
	repository.orders = repository.FindAllOrder()
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

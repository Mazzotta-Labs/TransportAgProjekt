package mysql

import (
	"TransportAgProjekt1/model/entity"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type MySqlRepository struct {
	drivers []entity.Driver
}

func NewMySqlRepository() *MySqlRepository {
	repository := MySqlRepository{}
	db = openDatabase()
	repository.drivers = repository.FindAllDriver()
	return &repository
}

func (r *MySqlRepository) FindAllDriver() []entity.Driver {
	var drivers []entity.Driver
	result, err := db.Query("select * from Driver join Vehicle using (id)")
	fmt.Println(result)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var driver entity.Driver

		err := result.Scan(&driver.DriverId, &driver.Name, &driver.Prename, &driver.VehicleId, &driver.Brand, &driver.Number)
		if err != nil {
			panic(err.Error())
		}
		drivers = append(drivers, driver)
	}
	return drivers
}

func (r *MySqlRepository) AddDriver(driver entity.Driver) {
	driverId := driver.DriverId
	name := driver.Name
	prename := driver.Prename
	vehicleId := driver.VehicleId
	brand := driver.Brand
	number := driver.Number

	stmt, err := db.Prepare("insert into model values (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(driverId, name, prename, vehicleId, brand, number)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateDriver(driver entity.Driver) {
	//TODO
}

func (r *MySqlRepository) DeleteDriver(driver entity.Driver) {
	//TODO
}

/*
type MySqlRepository struct {
	books []entity.entity
}

func NewMySqlRepository() * MySqlRepository{
	repository := MySqlRepository{}
	db = openDatabase()
	repository.books = repository.FindAll()
	return &repository
}

func (r *MySqlRepository) FindAll() []entity.Book{
	var books []entity.Book
	result, err := db.Query("select * from model")
	fmt.Println(result)
	if err != nil{
		panic(err.Error())
	}

	for result.Next(){
		var book entity.Book

		err := result.Scan(&book.ISBN, &book.Title, &book.Author, &book.PublishedYear, &book.Borrowed)
		if err != nil{
			panic(err.Error())
		}
		books = append(books, book)
	}
	return books
}

func (r *MySqlRepository) AddBook(book entity.Book){
	isbn := book.ISBN
	title := book.Title
	author := book.Author
	year := book.PublishedYear
	borrowed := book.Borrowed

	stmt, err := db.Prepare("insert into model values (?,?,?,?,?)")
	if err != nil{
		panic(err.Error())
	}
	_, err = stmt.Exec(isbn,title,author,year,borrowed)
	if err != nil{
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateBook(book entity.Book) {
	stmt, err := db.Prepare("update model set borrowed = ? where isbn = ?")
	if err != nil{
		panic(err.Error())
	}
	_, err = stmt.Exec(book.Borrowed,book.ISBN)
	if err != nil{
		panic(err.Error())
	}
}
*/

func openDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:Passwort@tcp(localhost:3306)/transportag")
	if err != nil {
		panic(err.Error())
	}
	return db
}

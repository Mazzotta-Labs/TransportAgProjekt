package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mvc-booklibrary/book/entity"
)

var db *sql.DB

type MySqlRepository struct {
	books []entity.Book
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


func openDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:Gamer-zone22@tcp(localhost:3306)/mybooks")
	if err != nil{
		panic(err.Error())
	}
	return db
}


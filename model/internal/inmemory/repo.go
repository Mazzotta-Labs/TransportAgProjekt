package inmemory

import "TransportAgProjekt1/model/entity"

type DriverInMemoryRepository struct {
	drivers []entity.Driver
}

func (r *DriverInMemoryRepository) FindAll() []entity.Driver {
	return r.drivers
}

func (r *DriverInMemoryRepository) AddBook(driver entity.Driver) {
	r.drivers = append(r.drivers, driver)
}

func (r *DriverInMemoryRepository) UpdateBook(driver entity.Driver) {
	for i, d := range r.drivers {
		if d.DriverId == driver.DriverId {
			r.drivers[i] = driver
		}
	}
}

func (r *DriverInMemoryRepository) DeleteDriver(driver entity.Driver) {
	//TODO
}

/*
type BookInMemoryRepository struct {
	books []entity.Book
}

func (r *BookInMemoryRepository) FindAll() []entity.Book  {
	return r.books
}

func (r *BookInMemoryRepository) AddBook(book entity.Book)  {
	r.books = append(r.books, book)
}

func (r *BookInMemoryRepository) UpdateBook(book entity.Book)  {
	for i, b := range r.books {
		if b.ISBN == book.ISBN {
			r.books[i] = book
		}
	}
}

*/

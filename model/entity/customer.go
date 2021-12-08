package entity

// Customer as type
type Customer struct {
	customerId      int
	customerName    string
	customerPrename string
	addressId       int
	telNr           string
}

// Address as type
type Address struct {
	addressId int
	street    string
	plz       string
	town      string
	country   string
}

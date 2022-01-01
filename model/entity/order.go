package entity

// Order as type
type Order struct {
	OrderId    int
	ProductsId []int
	ProductId  int
	OrderDate  string
	CustomerId int
	DriverId   int
}

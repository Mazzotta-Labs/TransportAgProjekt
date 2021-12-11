package entity

// Order as type
type Order struct {
	OrderId    int
	ProductsId []int
	OrderDate  int //passendes Zeitformat?
	CustomerId int
	DriverId   int
}

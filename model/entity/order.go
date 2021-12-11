package entity

import "time"

// Order as type
type Order struct {
	OrderId    int
	ProductsId []int
	OrderDate  time.Time //passendes Zeitformat?
	CustomerId int
	DriverId   int
}

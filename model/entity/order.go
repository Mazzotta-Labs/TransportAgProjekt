package entity

import "time"

// Order as type
type Order struct {
	OrderId    int
	ProductsId int
	OrderDate  time.Time
	CustomerId int
	DriverId   int
}

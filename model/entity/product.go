package entity

// Product as type
type Product struct {
	ProductId   int
	Description string
	Name        string
	CategoryId  int
}

// Category as type
type Category struct {
	CategoryId int
	Weight     float64
	Size       float64
	Fragile    bool
	Name       string
}

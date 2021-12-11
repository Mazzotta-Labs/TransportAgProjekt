package entity

// Product as type
type Product struct {
	ProductId   int
	Description string
	Name        string
	categoryId  int
}

// Categorie as type
type Categorie struct {
	CategoryId int
	Weight     float64
	Size       float64
	Fragile    bool
	Name       string
}

package entity

// Product as type
type Product struct {
	productId   int
	description string
	name        string
	categoryId  int
}

// Categorie as type
type Categorie struct {
	categoryId int
	weight     float64
	size       float64
	fragile    bool
	name       string
}

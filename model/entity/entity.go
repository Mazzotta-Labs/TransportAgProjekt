package entity

// Order as type
type Order struct {
	orderId		int
	products	[]int
	orderDate	int
	customerId	int
	driverId	int
}

// Product as type
type Product struct {
	productId	int
	description string
	categoryId	int
}

// Categorie as type
type Categorie struct {
	categoryId	int
	weight		float64
	size		float64
	fragile		bool
	name		string
}

// Driver as type
type Driver struct {
	driverId	int
	name		string
	prename		string
	vehicleId	int
}

// Vehicle as type
type Vehicle struct {
	vehicleId	int
	brand		string
	number		string
}

// Customer as type
type Customer struct {
	customerId		int
	customerName 	string
	customerPrename	string
	addressId		int
	telNr			string
}

// Address as type
type Address struct {
	addressId	int
	street		string
	plz			string
	town		string
	country		string
}
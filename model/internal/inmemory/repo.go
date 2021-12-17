package inmemory

import "TransportAgProjekt/model/entity"

type InMemoryRepository struct {
	drivers   []entity.Driver
	customers []entity.Customer
	orders    []entity.Order
	products  []entity.Product
}

//
//Customer
//
func (r *InMemoryRepository) FindAllCustomer() []entity.Customer {
	return r.customers
}

func (r *InMemoryRepository) FindCustomer() []entity.Customer {
	return r.customers
}

func (r *InMemoryRepository) AddCustomer(customer entity.Customer) {
	r.customers = append(r.customers, customer)
}

func (r *InMemoryRepository) UpdateCustomer(customer entity.Customer) {
	for i, d := range r.customers {
		if d.CustomerId == customer.CustomerId {
			r.customers[i] = customer
		}
	}
}

func (r *InMemoryRepository) DeleteCustomer(customer entity.Customer) {
	for i, d := range r.customers {
		if d.CustomerId == customer.CustomerId {
			r.customers[i] = r.customers[len(r.customers)-1]
		}
	}
}

//
//Driver
//
func (r *InMemoryRepository) FindAllDriver() []entity.Driver {
	return r.drivers
}

func (r *InMemoryRepository) AddDriver(driver entity.Driver) {
	r.drivers = append(r.drivers, driver)
}

func (r *InMemoryRepository) UpdateDriver(driver entity.Driver) {
	for i, d := range r.drivers {
		if d.DriverId == driver.DriverId {
			r.drivers[i] = driver
		}
	}
}

func (r *InMemoryRepository) DeleteDriver(driver entity.Driver) {
	//TODO
}

//
//Order
//
func (r *InMemoryRepository) FindAllOrder() []entity.Order {
	return r.orders
}

func (r *InMemoryRepository) AddOrder(order entity.Order) {
	r.orders = append(r.orders, order)
}

func (r *InMemoryRepository) UpdateOrder(order entity.Order) {
	for i, d := range r.orders {
		if d.OrderId == order.OrderId {
			r.orders[i] = order
		}
	}
}

func (r *InMemoryRepository) DeleteOrder(order entity.Order) {
	for i, d := range r.orders {
		if d.OrderId == order.OrderId {
			r.orders[i] = r.orders[len(r.orders)-1]
		}
	}
}

//
//Product
//
// func (r *InMemoryRepository) FindAllProduct() []entity.Product {
// 	return r.products
// }

// func (r *InMemoryRepository) AddProduct(product entity.Product) {
// 	r.products = append(r.products, product)
// }

// func (r *InMemoryRepository) UpdateProduct(product entity.Product) {
// 	for i, d := range r.products {
// 		if d.ProductId == product.ProductId {
// 			r.products[i] = product
// 		}
// 	}
// }

// func (r *InMemoryRepository) DeleteProduct(product entity.Product) {
// 	for i, d := range r.products {
// 		if d.ProductId == product.ProductId {
// 			r.products[i] = r.products[len(r.products)-1]
// 		}
// 	}
// }

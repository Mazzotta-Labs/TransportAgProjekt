package inmemory

import "TransportAgProjekt/model/entity"

type InMemoryRepository struct {
	drivers   []entity.Driver
	customers []entity.Customer
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

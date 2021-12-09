package mysql

import (
	"TransportAgProjekt/model/entity"
	"strconv"
)

func (r *MySqlRepository) FindAllCustomer() []entity.Customer {
	var customers []entity.Customer
	result, err := db.Query("select c.id, c.name, c.prename, c.tel_nr, a.id, a.street, a.plz, a.town, a.country from customer c left join address a on a.id = c.address_id  ")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var customer entity.Customer

		err := result.Scan(&customer.CustomerId, &customer.CustomerName, &customer.CustomerPrename, &customer.TelNr, &customer.AddressId, &customer.Street, &customer.Plz, &customer.Town, &customer.Country)
		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer)
	}
	return customers
}

func (r *MySqlRepository) AddCustomer(customer entity.Customer) {
	customerId := customer.CustomerId
	customerName := customer.CustomerName
	customerPrename := customer.CustomerPrename
	telNr := customer.TelNr
	addressId := customer.AddressId
	addressId = customerId //sollte gleich sein wie customer id
	street := customer.Street
	plz := customer.Plz
	town := customer.Town
	country := customer.Country

	stmt, err := db.Prepare("insert into Address values (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(addressId, street, plz, town, country) //brauche CustomerId gleichzeitig als AddressID da 1:1 Beziehung
	if err != nil {
		panic(err.Error())
	}

	stmt, err = db.Prepare("insert into Customer values (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(customerId, customerName, customerPrename, telNr, addressId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateCustomer(customer entity.Customer) {
	customerId := customer.CustomerId
	customerName := customer.CustomerName
	customerPrename := customer.CustomerPrename
	telNr := customer.TelNr
	addressId := customer.AddressId
	addressId = customerId //sollte gleich sein wie customer id
	street := customer.Street
	plz := customer.Plz
	town := customer.Town
	country := customer.Country

	stmt, err := db.Prepare("update address set street = ?, plz = ?, town = ?, country = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(street, plz, town, country, addressId) //brauche CustomerId gleichzeitig als AddressID da 1:1 Beziehung
	if err != nil {
		panic(err.Error())
	}

	stmt, err = db.Prepare("update customer set name = ?, prename = ?, tel_nr = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(customerName, customerPrename, telNr, customerId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) DeleteCustomer(customer string) {
	customerId, _ := strconv.Atoi(customer)
	addressId := customerId
	stmt, err := db.Prepare("delete from customer where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(addressId) //brauche CustomerId gleichzeitig als AddressID da 1:1 Beziehung
	if err != nil {
		panic(err.Error())
	}

	stmt, err = db.Prepare("delete from address where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(customerId)
	if err != nil {
		panic(err.Error())
	}
}

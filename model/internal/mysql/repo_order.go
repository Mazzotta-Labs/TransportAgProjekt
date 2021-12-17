package mysql

import (
	"TransportAgProjekt/model/entity"
)

func (r *MySqlRepository) FindAllOrder() []entity.Order {
	var orders []entity.Order
	result, err := db.Query("select o.id, o.order_date, c.id, c.name, c.prename, c.tel_nr, a.id, a.street, a.plz, a.town, a.country, d.id, d.prename, d.name, v.id, v.brand, v.number from `order` o left join customer c on c.id = o.customer_id left join driver d on d.id = o.driver_id left join address a on c.address_id = a.id left join vehicle v on d.vehicle_id = v.id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var order entity.Order

		err := result.Scan(&order.CustomerId, &order.DriverId, &order.OrderDate, &order.OrderId, &order.ProductsId)
		if err != nil {
			panic(err.Error())
		}
		orders = append(orders, order)
	}
	return orders
}

func (r *MySqlRepository) AddOrder(order entity.Order) {
	customerId := order.CustomerId
	driverId := order.DriverId
	orderDate := order.OrderDate
	orderId := order.OrderId
	productId := order.ProductsId

	stmt, err := db.Prepare("insert into model values (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(customerId, driverId, orderDate, orderId, productId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateOrder(order entity.Order) {
	// customerId := order.CustomerId
	// driverId := order.DriverId
	// orderDate := order.OrderDate
	// orderId := order.OrderId
	// productId := order.ProductsId

}

func (r *MySqlRepository) DeleteOrder(order entity.Order) {
	// orderId, _ := strconv.Atoi(order)
}

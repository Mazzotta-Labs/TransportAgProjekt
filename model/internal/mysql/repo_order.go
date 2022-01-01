package mysql

import (
	"TransportAgProjekt/model/entity"
)

func (r *MySqlRepository) FindAllOrder() []entity.Order {
	var orders []entity.Order
	result, err := db.Query("select o.id, o.order_date, c.id, d.id, ifnull(p.id,0) from `order` o left join customer c on c.id = o.customer_id left join driver d on d.id = o.driver_id left join ordertoproduct op on op.order_id = o.id left join product p on p.id = op.product_id order by o.id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var order entity.Order

		err := result.Scan(&order.OrderId, &order.OrderDate, &order.CustomerId, &order.DriverId, &order.ProductId)
		if err != nil {
			panic(err.Error())
		}
		orders = append(orders, order)
	}
	return orders
}

func (r *MySqlRepository) AddOrder(order entity.Order) {
	orderId := order.OrderId
	customerId := order.CustomerId
	driverId := order.DriverId
	orderDate := order.OrderDate
	products := order.ProductsId

	stmt, err := db.Prepare("insert into `Order` (`order_date`,`customer_id`,`driver_id`) values (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(orderDate, customerId, driverId)
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("select LAST_INSERT_ID()")
	for result.Next() {
		err := result.Scan(&orderId)
		if err != nil {
			panic(err.Error())
		}
	}

	for i := range products {
		stmt, err := db.Prepare("insert into `ordertoproduct` (`order_id`,`product_id`) values (?,?)")
		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(orderId, products[i])
		if err != nil {
			panic(err.Error())
		}
	}
}

func (r *MySqlRepository) UpdateOrder(order entity.Order) {
	stmt, err := db.Prepare("update `order` set customer_id = ?, driver_id = ?, order_date = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(order.CustomerId, order.DriverId, order.OrderDate, order.OrderId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) DeleteOrder(orderId string) {
	stmt, err := db.Prepare("delete from OrderToProduct where order_id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(orderId) //Order wird wegen Delete Cascade auch gelöscht
	if err != nil {
		panic(err.Error())
	}
}

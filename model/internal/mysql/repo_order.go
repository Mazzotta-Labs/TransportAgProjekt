package mysql

import (
	"TransportAgProjekt/model/entity"
	"time"
)

func (r *MySqlRepository) FindAllOrder() []entity.Order {
	var orders []entity.Order
	result, err := db.Query("select o.id, o.order_date, c.name, c.prename, a.street, a.plz, a.town, a.country, d.name, d.prename, v.number from `order` o left join customer c on c.id = o.customer_id left join address a on c.address_id = a.id left join driver d on d.id = o.driver_id left join vehicle v on d.vehicle_id = v.id left join ordertoproduct op on op.order_id = o.id left join product p on p.id = op.product_id order by o.id")
	if err != nil {
		println("query error")
		time.Sleep(10 * time.Second)
		panic(err.Error())
	}

	for result.Next() {
		var order entity.Order

		err := result.Scan(&order.OrderId, &order.OrderDate, &order.CustomerName, &order.CustomerPrename, &order.Street, &order.Plz, &order.Town, &order.Country, &order.Name, &order.Prename, &order.Number)
		if err != nil {
			panic(err.Error())
		}
		orders = append(orders, order)
	}
	return orders
}

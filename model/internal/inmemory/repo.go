package inmemory

import "TransportAgProjekt/model/entity"

type InMemoryRepository struct {
	orders []entity.Order
}

//
//Order
//
func (r *InMemoryRepository) FindAllOrder() []entity.Order {
	return r.orders
}

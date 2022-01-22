package internal

import (
	"TransportAgProjekt/model/entity"
	"TransportAgProjekt/model/internal/mysql"
)

var R = mysql.NewMySqlRepository()

type Repository interface {
	FindAllOrder() []entity.Order
}

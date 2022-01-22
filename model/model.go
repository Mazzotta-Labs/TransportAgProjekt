package model

import (
	"TransportAgProjekt/model/entity"
	"TransportAgProjekt/model/internal"
	"os"
)

func FindAllOrder() []entity.Order {
	return internal.R.FindAllOrder()
}

// Model
func ShutDown() {
	os.Exit(0)
}

package usecase

import (
	"github.com/QUDUSKUNLE/microservices/services/order-service/protogen/golang/orders"
)

// GetByEmail implements v1.RepoInterface.
func (use *UseCase) AddOrder(order *orders.Order) error {
	return use.repo.AddOrder(order)
}

func (use *UseCase) Get(id string) (*orders.Order, error) {
	return use.repo.Get(id)
}

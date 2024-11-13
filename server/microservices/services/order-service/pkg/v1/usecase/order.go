package usecase

import (
	"github.com/QUDUSKUNLE/microservices/services/order-service/protogen/golang/orders"
)

// GetByEmail implements v1.RepoInterface.
func (use *UseCase) AddOrder(order *orders.Order) error {
	return use.repo.AddOrder(order)
}

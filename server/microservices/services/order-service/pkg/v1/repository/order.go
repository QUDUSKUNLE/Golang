package repo

import (
	"github.com/QUDUSKUNLE/microservices/services/order-service/protogen/golang/orders"
)

// Create implements v1.RepoInterface.
func (repository *Repository) AddOrder(order *orders.Order) error {
	return repository.database.Create(&order).Error
}

package repo

import (
	"github.com/QUDUSKUNLE/microservices/order-service/protogen/golang/orders"
)

// Create implements v1.RepoInterface.
func (repository *Repository) AddOrder(order *orders.Order) error {
	return repository.database.Create(&order).Error
}

func (repository *Repository) Get(id string) (*orders.Order, error) {
	var order *orders.Order
	err := repository.database.Where("id = ?", id).First(order).Error
	return order, err
}

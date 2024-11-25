package v1

import (
	"github.com/QUDUSKUNLE/microservices/order-service/protogen/golang/orders"
)

type RepositoryInterface interface {
	AddOrder(order *orders.Order) error
	Get(id string) (*orders.Order, error)
}

type UseCaseInterface interface {
	AddOrder(order *orders.Order) error
	Get(id string) (*orders.Order, error)
}

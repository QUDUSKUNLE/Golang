package models

import (
	"github.com/QUDUSKUNLE/microservices/services/order-service/protogen/golang/orders"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	collection []*orders.Order
}

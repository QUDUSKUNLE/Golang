package models

import (
	"github.com/QUDUSKUNLE/microservices/order-service/protogen/golang/orders"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	collection []*orders.Order
}

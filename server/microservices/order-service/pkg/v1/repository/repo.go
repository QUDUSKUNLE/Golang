package repo

import (
	interfaces "github.com/QUDUSKUNLE/microservices/order-service/pkg/v1"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) interfaces.RepositoryInterface {
	return &Repository{database: database}
}

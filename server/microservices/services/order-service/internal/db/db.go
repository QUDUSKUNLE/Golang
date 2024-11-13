package database

import (
	"fmt"
	"log"

	"github.com/QUDUSKUNLE/microservices/services/order-service/internal/models"
	interfaces "github.com/QUDUSKUNLE/microservices/services/order-service/pkg/v1"
	repo "github.com/QUDUSKUNLE/microservices/services/order-service/pkg/v1/repository"
	usecase "github.com/QUDUSKUNLE/microservices/services/order-service/pkg/v1/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn(host, user, dbname, password string) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=5432 sslmode=disable", host, user, dbname, password)), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	if err = db.AutoMigrate(&models.Order{}); err != nil {
		fmt.Println(err)
	}
	return db
}

func InitUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	userRepo := repo.NewRepository(db)
	return usecase.New(userRepo)
}

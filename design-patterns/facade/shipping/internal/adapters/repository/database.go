package repository

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func PostgresSQLConnection() (*PostgresRepository, error) {
	point := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(point, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int.")
	}
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	DB, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database: %w", err)
	}
	if err := DB.AutoMigrate(
		&domain.User{},
		&domain.Shipping{},
		&domain.PickUp{},
	); err != nil {
		log.Fatalf("Error running struct model migration: %s", err.Error())
	}

	postgresDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	// Set database connection settings.
	postgresDB.SetMaxOpenConns(maxConn)
	postgresDB.SetMaxIdleConns(maxIdleConn)
	postgresDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Ping PostgresRepository
	if err = postgresDB.Ping(); err != nil {
		return nil, fmt.Errorf("error, not send ping to database: %w", err)
	}
	return &PostgresRepository{
		db: DB }, nil
}

func OpenDBConnection() (*PostgresRepository, error) {
	DB, err := PostgresSQLConnection()
	if err != nil {
		return nil, err
	}
	return DB, nil
}

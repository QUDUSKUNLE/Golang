package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/QUDUSKUNLE/shipping/src/database/repository"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	*repository.Database
}

func PostgresSQLConnection() (*gorm.DB, error) {
	point := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(point, 10, 32);
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
		&model.User{},
		&model.Shipping{},
		&model.PickUp{},
	); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}


	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	// Set database connection settings.
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Ping Database
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error, not send ping to database: %w", err)
	}
	return DB, nil
}

func OpenDBConnection() (*PostgresRepository, error) {
	db, err := PostgresSQLConnection()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		Database: &repository.Database{ DB: db },
	}, nil
}

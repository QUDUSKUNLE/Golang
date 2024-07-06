package database

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"github.com/QUDUSKUNLE/shipping/src/database/repository"
	"github.com/jmoiron/sqlx"

  _ "github.com/jackc/pgx/v4/stdlib" 
	_ "github.com/lib/pq"
)

type ShippingDB struct {
	*repository.Database
}

func PostgresSQLConnection() (*sqlx.DB, error) {
	point := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(point, 10, 32);
	if err != nil {
		fmt.Println("Error parsing str to int.")
	}
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
  maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
  maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	DB, err := sqlx.Connect("pgx", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database: %w", err)
	}

	// Set database connection settings.
	DB.SetMaxOpenConns(maxConn)
	DB.SetMaxIdleConns(maxIdleConn)
	DB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Ping Database
	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("error, not send ping to database: %w", err)
	}
	return DB, nil
}

func OpenDBConnection() (*ShippingDB, error) {
	db, err := PostgresSQLConnection()
	if err != nil {
		return nil, err
	}
	return &ShippingDB{
		Database: &repository.Database{ DB: db },
	}, nil
}

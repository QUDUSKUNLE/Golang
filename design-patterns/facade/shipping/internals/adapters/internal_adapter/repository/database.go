package repository

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries: true,
			Colorful: true,
		},
	)
	DB, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database: %w", err)
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

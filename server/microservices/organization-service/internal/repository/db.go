package repository

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/services"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresRepository struct {
	db *gorm.DB
}

// CreateShipping implements ports.RepositoryPorts.
func (database *PostgresRepository) CreateShipping(shipping []*domain.Shipping) error {
	panic("unimplemented")
}

// GetShipping implements ports.RepositoryPorts.
func (database *PostgresRepository) GetShipping(shippingID uuid.UUID, userID uuid.UUID) (*domain.Shipping, error) {
	panic("unimplemented")
}

// GetShippings implements ports.RepositoryPorts.
func (database *PostgresRepository) GetShippings(shippingID uuid.UUID, status string) ([]*domain.Shipping, error) {
	panic("unimplemented")
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
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	DB, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"))), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 newLogger,
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
		db: DB}, nil
}

func OpenDBConnection() (*gorm.DB, error) {
	DB, err := PostgresSQLConnection()
	if err != nil {
		return nil, err
	}
	return DB.db, nil
}

func InitUserServer(db *gorm.DB) ports.ShippingPorts {
	shippingRepo := NewRepository(db)
	return services.InternalServicesAdapter(shippingRepo)
}

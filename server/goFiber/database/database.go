package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"strconv"
	"github.com/QUDUSKUNLE/gofiber/config"
	"github.com/QUDUSKUNLE/gofiber/database/queries"
	"github.com/jmoiron/sqlx"

  _ "github.com/jackc/pgx/v4/stdlib" 
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Queries struct {
	*queries.Fiber
}

func PostgresSQLConnection() (*sqlx.DB, error) {
	point := config.Config("DB_PORT")
	port, err := strconv.ParseUint(point, 10, 32);
	if err != nil {
		fmt.Println("Error parsing str to int.")
	}
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
  maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
  maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	DB, err := sqlx.Connect("pgx", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database: %w", err)
	}

	// Set database connection settings.
	DB.SetMaxOpenConns(maxConn)
	DB.SetMaxIdleConns(maxIdleConn)
	DB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Ping Database
	if err = DB.Ping(); err != nil {
		defer DB.Close(); // Close database connection
		return nil, fmt.Errorf("error, not send ping to database: %w", err)
	}
	return DB, nil
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgresSQLConnection()
	if err != nil {
		return nil, err
	}
	return &Queries{
		Fiber: &queries.Fiber{DB: db},
	}, nil
}

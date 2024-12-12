package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DatabaseConnection() *Queries {
	database, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return New(database)
}

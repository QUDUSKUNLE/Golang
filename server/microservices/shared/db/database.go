package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DatabaseConnection(DB_URL string) *Queries {
	database, err := pgxpool.New(context.Background(), DB_URL)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return New(database)
}

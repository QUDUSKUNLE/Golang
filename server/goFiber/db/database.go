package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"gofiber/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	point := config.Config("DB_PORT")
	port, err := strconv.ParseUint(point, 10, 32);
	if err != nil {
		fmt.Println("Error parsing str to int.")
	}
	DB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Connection opened to database.")
	return nil
}

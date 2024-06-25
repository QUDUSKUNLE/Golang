package handlers

import (
	"strconv"
	"gofiber/database"
	"github.com/gofiber/fiber/v2"
)

func HandlerHelper(c *fiber.Ctx) (*database.Queries, int, error) {
	ID := c.Params("id")
	id, err := strconv.Atoi(ID);
	if err != nil {
		panic(err)
	}
	// Create a database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return nil, 0, err
	}
	return db, id, nil
}

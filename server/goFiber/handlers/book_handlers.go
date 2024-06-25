package handlers

import (
	"gofiber/database"
	"gofiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(context *fiber.Ctx) error {
	// Create a database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	books, err := db.QueryGetBooks();
	if err != nil {
		defer db.Close();
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err.Error(),
		})
	}

	defer db.Close(); // close database connection
	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"books": books,
		"message": "All books returned successfully.",
	})
}

func CreateBook(context *fiber.Ctx) error {
	book := &models.Book{}

	if err := context.BodyParser(book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Create a database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := db.QueryCreateBook(book); err != nil {
		defer db.Close() // Close database
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	
	defer db.Close() // Close database
	return context.JSON(&fiber.Map{
		"success": true,
		"message": "Book created successfully.",
		"book": book,
	})
}

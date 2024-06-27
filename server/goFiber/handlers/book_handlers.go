package handlers

import (
	"fmt"
	"gofiber/database"
	"gofiber/models"

	"gofiber/utils"

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
	defer db.Close()
	books, err := db.QueryGetBooks();
	if err != nil {
		db.Close();
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err.Error(),
		})
	}

	db.Close(); // close database connection
	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"books": books,
		"message": "All books returned successfully.",
	})
}

func CreateBook(context *fiber.Ctx) error {
	bookAttr := &models.BookAttrs{}
	book := new(models.Book)
	val, err := models.BookAttrs.Value(models.BookAttrs{});
	if err != nil {
		panic(err)
	}
	if err := context.BodyParser(&book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	bookAttr.Scan(val)
	fmt.Println(book, "YYYYY")

	// Create a database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	defer db.Close()
	newProduct := models.NewBook()
	book.ID = newProduct.ID
	book.CreatedAt = newProduct.CreatedAt
	book.UpdatedAt = newProduct.UpdatedAt
	book.BookAttrs = models.BookAttrs{
		Picture: book.BookAttrs.Picture,
		Description: book.BookAttrs.Description,
		Rating: book.BookAttrs.Rating,
	}

	// Create a new validator
	validate := utils.Validator()
	if err := validate.Struct(book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": utils.ValidatorErrors(err),
		})
	}
	if err := db.QueryCreateBook(book); err != nil {
		db.Close() // Close database
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// fmt.Println(book)
	db.Close() // Close database
	return context.JSON(&fiber.Map{
		"success": true,
		"message": "Book created successfully.",
		"book": book,
	})
}

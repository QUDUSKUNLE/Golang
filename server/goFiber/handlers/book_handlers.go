package handlers

import (
	"fmt"
	"gofiber/database"
	"gofiber/models"

	"gofiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func GetBook(context *fiber.Ctx) error {
	// Catch book ID from URL
	id, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}

	// Create a database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}

	// Get book by ID
	book, err := db.QueryGetBook(id)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}
	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"error": false,
		"book": book,
	})
}

func DeleteBook(context *fiber.Ctx) error {
	// Create new Book struct
	book := &models.Book{}
	// Check, if received JSON data is valid.
	if err := context.BodyParser(book); err != nil {
		// Return status 400 and error message.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"message": err.Error(),
		})
	}

	// Create a new validator for a Book model.
	validate := utils.Validator()

	// Validate only one book field ID.
	if err := validate.StructPartial(book, "id"); err != nil {
			// Return, if some fields are not valid.
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": true,
					"message": utils.ValidatorErrors(err),
			})
	}
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}

	// Checking, if book with given ID is exists.
	foundedBook, err := db.QueryGetBook(book.ID)
	if err != nil {
		// Return status 404 and book not found error.
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"message": "book with this ID not found",
		})
	}

	// Delete book by given ID.
	if err := db.QueryDeleteBook(foundedBook.ID); err != nil {
		// Return status 500 and error message.
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}
	// Return status 204 no content.
	return context.SendStatus(fiber.StatusNoContent)
}

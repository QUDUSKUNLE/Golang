package handlers

import (
	"github.com/QUDUSKUNLE/gofiber/src/database"
	"github.com/QUDUSKUNLE/gofiber/src/models"
	"github.com/QUDUSKUNLE/gofiber/src/utils"
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
	books, err := db.QueryGetBooks();
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"books": books,
		"message": "All books returned successfully.",
	})
}

func CreateBook(context *fiber.Ctx) error {
	// bookAttr := new(models.BookAttrs)
	book := new(models.Book)

	if err := context.BodyParser(&book); err != nil {
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

	newProduct := models.NewBook()
	book = &models.Book{
		ID: newProduct.ID,
		CreatedAt: newProduct.CreatedAt,
		UpdatedAt: newProduct.UpdatedAt,
		BookAttrs: models.BookAttrs{
			Picture: book.BookAttrs.Picture,
			Description: book.BookAttrs.Description,
			Rating: book.BookAttrs.Rating,
		},
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
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

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
	book := new(models.Book)
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

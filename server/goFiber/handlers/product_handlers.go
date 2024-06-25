package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/database"
	"gofiber/models"
	"strconv"
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

// Home
func FiberHome(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).SendString("Fiber framework!")
}

// Get All Products from database
func GetAllProducts(context *fiber.Ctx) error {
	// Create a database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	products, err := db.QueryGetProducts();
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
		"products": products,
		"message": "All product returned successfully",
	})
}

// Get Single Product from database
func GetSingleProduct(context *fiber.Ctx) error {
	db, id, err := HandlerHelper(context)
		if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// Fetch Product from the database
	product, err := db.QueryGetProduct(id)
	if err != nil {
		defer db.Close() // Close database connection
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	defer db.Close() // Close database connection
	return context.JSON(&fiber.Map{
		"success": true,
		"message": "Successfully fetched product",
		"product": &product,
	});
}

// Creating a Product
func CreateProduct(context *fiber.Ctx) error {
	product := &models.Product{}

	if err := context.BodyParser(product); err != nil {
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

	product.ID = models.NewProduct().ID
	if err := db.QueryCreateProduct(product); err != nil {
		defer db.Close() // Close database
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	
	defer db.Close() // Close database
	return context.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"product": product,
	})
}

// Delete a Product
func DeleteProduct(context *fiber.Ctx) error {
	db, id, err := HandlerHelper(context)
		if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := db.QueryDeleteProduct(id); err != nil {
		defer db.Close()
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
	}

	defer db.Close()
	return context.JSON(&fiber.Map{
		"success": true,
		"message": "Product deleted succcessfully",
	})
}

// Update a Product
func UpdateProduct(context *fiber.Ctx) error {
	product := &models.Product{}

	if err := context.BodyParser(product); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	db, id, err := HandlerHelper(context)
		if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := db.QueryUpdateProduct(id, product); err != nil {
		defer db.Close()
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
	}
	return context.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"success": true,
		"message": "Product updated successfully",
	})
}

package handlers

import (
	"encoding/json"
	"gofiber/database"
	"gofiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetHome(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).SendString("Hello, World!")
}

func GetBody(context *fiber.Ctx) error {
	body := context.Body()
	var result map[string]interface{}
	if err:= json.Unmarshal(body, &result); err != nil {
		return context.Send(body)
	}
	return context.JSON(result)
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

	products, err := db.GetProducts();
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

// Get Singel Product from database
func GetSingleProduct(context *fiber.Ctx) error {
	ID := context.Params("id")
	id, err := strconv.Atoi(ID);
	if err != nil {
		panic(err)
	}
	
	// Create a database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// Fetch Product from the database
	product, err := db.GetProduct(id)
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
	if err := db.CreateProduct(product); err != nil {
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

// Delete Product from DB
func DeleteProduct(context *fiber.Ctx) error {
	id := context.Params("id")
	_, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
		return nil
	}
	if err := context.JSON(&fiber.Map{
		"success": true,
		"message": "Product deleted succcessfully",
	}); err != nil {
		context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
		return nil
	}
	return nil
}

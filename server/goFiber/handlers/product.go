package handlers

import (
	"log"
	"database/sql"
	"encoding/json"
	"gofiber/db"
	"gofiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetHome(context *fiber.Ctx) error {
	return context.SendString("Hello, World!")
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
	rows, err := db.DB.Query("SELECT name, description, category, amount FROM products order by name")
	if err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
	}
	defer rows.Close()
	result := models.Products{}

	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount); err != nil {
			context.Status(500).JSON(&fiber.Map{
				"success": false,
				"error": err,
			})
			return nil
		}
		result.Products = append(result.Products, product)
	}
	// return result in JSON
	if err := context.Status(200).JSON(&fiber.Map{
		"success": true,
		"products": result,
		"message": "All product returned successfully",
	}); err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return nil
	}
	return nil
}

// Get Singel Product from database
func GetSingleProduct(context *fiber.Ctx) error {
	id := context.Params("id")
	product := models.Product{}

	row, err := db.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	defer row.Close();

	for row.Next() {
		switch err := row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category); err {
		case sql.ErrNoRows:
			log.Println("Now rows were returned!")
			context.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(product.Name, product.Amount, product.Category, product.Description)
		default:
			context.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}
	if err := context.JSON(&fiber.Map{
		"success": true,
		"message": "Successfully fetched product",
		"product": product,
	}); err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return nil
	}
	return nil
}

// Creating a Product
func CreateProduct(context *fiber.Ctx) error {
	product := new(models.Product)

	if err := context.BodyParser(product); err != nil {
		log.Println(err)
		context.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return nil
	}
	_, err := db.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)", product.Name, product.Description, product.Category, product.Amount)
	if err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return nil
	}
	
	if err := context.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"product": product,
	}); err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return nil
	}
	return nil
}

// Delete Product from DB
func DeleteProduct(context *fiber.Ctx) error {
	id := context.Params("id")
	_, err := db.DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
		return nil
	}
	if err := context.JSON(&fiber.Map{
		"success": true,
		"message": "Product deleted succcessfully",
	}); err != nil {
		context.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
		return nil
	}
	return nil
}

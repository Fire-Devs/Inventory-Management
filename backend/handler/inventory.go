package handler

import (
	"InventoryManagement/models"
	"InventoryManagement/repository"
	"github.com/gofiber/fiber/v3"
)

func CreateCategory(c fiber.Ctx) error {
	name, description := c.FormValue("name"), c.FormValue("description")
	if name == "" || description == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name or Description is missing",
		})
	}
	err := repository.InsertCategory(
		&models.Category{
			Name:        name,
			Description: description,
		})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Category created successfully"})
}

func FetchAllCategories(c fiber.Ctx) error {
	categories, err := repository.GetCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(categories)
}

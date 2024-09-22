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

func CreateSupplier(c fiber.Ctx) error {
	name, contactInfo := c.FormValue("name"), c.FormValue("contact_info")
	if name == "" || contactInfo == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name or Contact Info is missing",
		})
	}
	err := repository.InsertSupplier(
		&models.Supplier{
			Name:        name,
			ContactInfo: contactInfo,
		})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Supplier created successfully"})
}

func FetchAllSuppliers(c fiber.Ctx) error {
	suppliers, err := repository.GetSuppliers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(suppliers)
}

func CreateInventory(c fiber.Ctx) error {
	inventory := new(models.Inventory)
	if err := c.Bind().Body(inventory); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := repository.InsertInventory(inventory)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Inventory created successfully"})
}

func FetchAllInventory(c fiber.Ctx) error {
	inventory, err := repository.GetInventory()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(inventory)
}

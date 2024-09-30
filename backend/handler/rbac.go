package handler

import (
	"InventoryManagement/models"
	"github.com/gofiber/fiber/v3"
)

func AddPermissionToRoles(c fiber.Ctx) error {

	role := new(models.Role)
	if err := c.Bind().Body(role); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": role,
	})
}

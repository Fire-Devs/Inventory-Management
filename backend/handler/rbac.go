package handler

import (
	"InventoryManagement/models"
	"InventoryManagement/repository"
	"github.com/gofiber/fiber/v3"
)

func UpdateRoleToUser(c fiber.Ctx) error {
	role := new(models.RoleUser)
	if err := c.Bind().Body(role); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := repository.AssignRoleToUser(role.Email, role.RoleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Role assigned to user successfully",
	})
}

func InsertRoles(c fiber.Ctx) error {
	role := new(models.Role)
	if err := c.Bind().Body(role); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := repository.AddRoles(role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": role,
	})
}

func FetchRoles(c fiber.Ctx) error {
	roleName := c.Query("name")
	roles, err := repository.FetchRoles(roleName)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"roles": roles,
	})
}

func UpdateRoles(c fiber.Ctx) error {
	role := new(models.Role)
	if err := c.Bind().Body(role); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if role.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Role ID is required",
		})
	}

	err := repository.UpdateRoles(role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": role,
	})
}

func FetchAllPermissions(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"permissions": models.Permissions,
	})
}

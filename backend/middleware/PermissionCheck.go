package middleware

import (
	"InventoryManagement/handler"
	"InventoryManagement/repository"
	"github.com/gofiber/fiber/v3"
)

func CheckPermission(c fiber.Ctx, permission string) error {
	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	err := handler.ParseJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// get the role of the user from the token

	role, err := repository.FetchPermissionfromUser(permission)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"permission": role,
	})

	return c.Next()

}

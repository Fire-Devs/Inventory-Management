package middleware

import (
	"InventoryManagement/handler"
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

	return c.Next()

}

package middleware

import (
	"InventoryManagement/handler"
	"github.com/gofiber/fiber/v3"
)

func IsAuthorized(c fiber.Ctx) (err error) {
	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})

	}

	err2 := handler.ParseJWT(cookie)
	if err2 != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})

	}

	return c.Next()

}

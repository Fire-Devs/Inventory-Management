package middleware

import (
	"InventoryManagement/config"
	"InventoryManagement/repository"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func CheckPermission(permission string) fiber.Handler {
	return func(c fiber.Ctx) error {
		cookie := c.Cookies("token")
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		conf := config.LoadConfig()
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Jwt.Secret), nil
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		data := claims["email"].(string)

		role, err := repository.FetchPermissionfromUser(data)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		for _, v := range role {
			if v == permission {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
}

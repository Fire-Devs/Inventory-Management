package routes

import (
	"InventoryManagement/handler"
	"InventoryManagement/middleware"
	"github.com/gofiber/fiber/v3"
)

func HandleRoutes(app *fiber.App) {

	auth := app.Group("/auth")

	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
	auth.Get("/verify", handler.VerifyToken)

	app.Get("/user", GetUser, middleware.IsAuthorized)
}

func GetUser(c fiber.Ctx) error {

	return c.Send([]byte("Hello, World!"))
}

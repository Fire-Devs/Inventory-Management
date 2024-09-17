package routes

import (
	"InventoryManagement/handler"
	"github.com/gofiber/fiber/v3"
)

func HandleRoutes(app *fiber.App) {

	auth := app.Group("/auth")

	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
	auth.Get("/verify", handler.VerifyToken)

	//app.Get("/user", GetUser, middleware.IsAuthorized)
}

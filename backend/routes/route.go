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

	app.Post("/categories", handler.CreateCategory, middleware.IsAuthorized)
	app.Get("/categories", handler.FetchAllCategories, middleware.IsAuthorized)
}

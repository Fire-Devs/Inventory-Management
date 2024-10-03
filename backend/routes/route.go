package routes

import (
	"InventoryManagement/handler"
	"InventoryManagement/middleware"
	"github.com/gofiber/fiber/v3"
)

func HandleRoutes(app *fiber.App) {

	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
	auth.Get("/verify", handler.VerifyToken)

	// RBAC routes
	app.Get("/permissions", handler.FetchAllPermissions, middleware.CheckPermission("read:role"))

	app.Post("/roles", handler.InsertRoles)
	app.Get("/roles", handler.FetchRoles)
	app.Put("/roles", handler.UpdateRoles)

	// Inventory addition routes
	app.Post("/categories", handler.CreateCategory, middleware.IsAuthorized)
	app.Post("/suppliers", handler.CreateSupplier, middleware.IsAuthorized)
	app.Post("/inventory", handler.CreateInventory, middleware.IsAuthorized)

	// Inventory fetch routes
	app.Get("/categories", handler.FetchAllCategories)
	app.Get("/suppliers", handler.FetchAllSuppliers)
	app.Get("/inventory", handler.FetchAllInventory)

}

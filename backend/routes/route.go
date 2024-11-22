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
	app.Put("/roleuser", handler.UpdateRoleToUser, middleware.CheckPermission("update:role"))

	app.Post("/roles", handler.InsertRoles, middleware.CheckPermission("create:role"))
	app.Get("/roles", handler.FetchRoles, middleware.CheckPermission("read:role"))
	app.Put("/roles", handler.UpdateRoles, middleware.CheckPermission("update:role"))

	// Inventory addition routes
	app.Post("/categories", handler.CreateCategory, middleware.CheckPermission("create:category"))
	app.Post("/suppliers", handler.CreateSupplier, middleware.CheckPermission("create:supplier"))
	app.Post("/inventory", handler.CreateInventory, middleware.CheckPermission("create:inventory"))

	// Inventory fetch routes
	app.Get("/categories", handler.FetchAllCategories)
	app.Get("/suppliers", handler.FetchAllSuppliers)
	app.Get("/inventory", handler.FetchAllInventory)

}

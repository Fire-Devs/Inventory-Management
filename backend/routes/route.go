package routes

import (
	"InventoryManagement/handler"
	"github.com/gofiber/fiber/v3"
)

func HandleRoutes(app *fiber.App) {

	// Inventory addition routes
	app.Post("/categories", handler.CreateCategory)
	app.Post("/suppliers", handler.CreateSupplier)
	app.Post("/inventory", handler.CreateInventory)
	app.Post("/inventory/price", handler.CreatePrice)

	// Inventory fetch routes
	app.Get("/categories", handler.FetchAllCategories)
	app.Get("/suppliers", handler.FetchAllSuppliers)
	app.Get("/inventory", handler.FetchAllInventory)

}

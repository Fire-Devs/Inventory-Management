package main

import (
	"InventoryManagement/config"
	"InventoryManagement/routes"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
		ServerHeader:    "Askar",
		AppName:         "Inventory Management",
	})
	app.Use(logger.New())

	conf := config.LoadConfig()

	routes.HandleRoutes(app)

	if err := app.Listen(conf.Server.Port); err != nil {
		log.Fatal(err)
	}
}

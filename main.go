package main

import (
	"IoT/database"
	"IoT/database/migration"
	"IoT/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Database initialization
	database.DBInit()

	// Database migration\
	migration.Migrate()

	// fiber app
	app := fiber.New()

	// Route initialization
	route.RouteInit(app)

	app.Listen(":3000")

}

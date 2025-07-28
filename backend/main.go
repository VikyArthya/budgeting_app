package main

import (
	"log"

	"github.com/VikyArthya/budgeting_app/database"
	"github.com/VikyArthya/budgeting_app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

    // Initialize a new Fiber app
    app := fiber.New()

	routes.Setup(app)

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
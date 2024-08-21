package main

import (
	"log"
	"trimmy/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Println("")
	log.Fatal(app.Listen(":8080"))
}

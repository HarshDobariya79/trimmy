package routes

import (
	"trimmy/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/u/:id", handlers.Redirect)
	app.Post("/shorten", handlers.ShortenURL)
}

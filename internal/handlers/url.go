package handlers

import (
	"trimmy/internal/services"
	"trimmy/pkg/config"

	"github.com/gofiber/fiber/v2"
)

// Process short URL redirect
func Redirect(c *fiber.Ctx) error {
	id := c.Params("id")
	originalURL, err := services.GetOriginalURL(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("URL not found")
	}
	return c.Redirect(originalURL, fiber.StatusMovedPermanently)
}

// Shorten an URL
func ShortenURL(c *fiber.Ctx) error {
	type request struct {
		URL string `json:"url"`
	}
	req := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	shortID, err := services.GenerateShortURL(req.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not shorten URL")
	}

	return c.JSON(fiber.Map{"shortUrl": config.Env["BACKEND_HOST"] + "/u/" + shortID})
}

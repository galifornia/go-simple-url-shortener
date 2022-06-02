package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", resolveURL)
	app.Post("/api/v1", shortenURL)
}

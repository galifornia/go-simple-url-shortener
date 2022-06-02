package main

import (
	"os"

	"github.com/go-delve/delve/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	app := fiber.New()

	app.Use(Logger.New())
	routes.SetupRoutes()
	app.Listen(os.Getenv("DOMAIN") + ":" + os.Getenv("APP_PORT"))
}

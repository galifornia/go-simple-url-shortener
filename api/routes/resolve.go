package routes

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func resolveURL(c *fiber.Ctx) error {
	body := new(types.Request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// implement rate limiting

	// check if input is a valid URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid URL"})
	}
	// check for domain error

	// enforce https

	return nil
}

package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func WithUserAuthorization(c *fiber.Ctx) error {
	log.Println("user id", c.Locals("user_id"))
	if c.Locals("user_id") == nil {
		return c.Status(401).Next()
	}

	return c.Next()
}

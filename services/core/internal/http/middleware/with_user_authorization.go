package middleware

import (
	"core/internal/db"
	"core/configs"
	
	"context"

	"github.com/gofiber/fiber/v2"
	"log"
)

func WithUserAuthorization(c *fiber.Ctx) error {
	if c.Locals("user_id") == "" {
		return ctx.Status(401).Next()
	} 

	return c.Next()
}
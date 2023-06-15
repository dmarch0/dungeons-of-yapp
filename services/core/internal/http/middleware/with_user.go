package middleware

import (
	"core/internal/db"
	"core/configs"
	
	"context"

	"github.com/gofiber/fiber/v2"
	"log"
)

func WithUser(c *fiber.Ctx) error {
	db_connection := db.GetDbConnection()
	config := configs.GetConfig()
	token := c.Cookies(config.HTTPServerConfig.AccessTokenCookieName)
	context := context.Background()

	if token != "" {
		userQuery := `
			SELECT users.id, email
			FROM user_tokens
			INNER JOIN users
			ON users.id = user_id
			WHERE token=$1
		`
		var id int
		var email string

		if err := db_connection.QueryRow(context, userQuery, token).Scan(&id, &email); err != nil {
			c.Locals("user_email", email)
			c.Locals("user_id", id)
		}
	}
	log.Println("token", token, config.HTTPServerConfig.AccessTokenCookieName)

	return c.Next()
}
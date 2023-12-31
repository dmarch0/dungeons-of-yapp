package routes

import "github.com/gofiber/fiber/v2"

func AddEndpoints(app *fiber.App) {
	apiRoot := app.Group("/api")
	v1 := apiRoot.Group("/v1")

	AddHelloEndpoints(v1)
	AddUserEndpoints(v1.Group("/user"))
	AddCharacterEndpoints(v1.Group("/character"))
}
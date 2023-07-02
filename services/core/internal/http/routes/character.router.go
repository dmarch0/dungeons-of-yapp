package routes

import (
	"core/internal/http/controllers"
	"core/internal/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddCharacterEndpoints(router fiber.Router) {
	router.Get("/", controllers.GetCharacterController)
	router.Post("/", controllers.PostCharacterController)
	router.Post(
		"/random",
		middleware.WithUser,
		middleware.WithUserAuthorization,
		controllers.PostCharacterRandomController,
	)
}

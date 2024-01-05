package routes

import (
	"auth/controllers"

	"github.com/gofiber/fiber/v2"
)

// userRoutes routes
func UserRoutes(app fiber.Router, uc *controllers.UsersController) {

	user := app.Group("/users")

	user.Post("/auth", func(c *fiber.Ctx) error {
		return uc.Authenticate(c)
	})

}

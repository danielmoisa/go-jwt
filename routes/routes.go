package routes

import (
	"github.com/danielmoisa/go-jwt/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Auth
	app.Post("/api/auth/register", controllers.Register)
	app.Post("api/auth/login", controllers.Login)
	app.Get("api/auth/user", controllers.GetUser)
	app.Post("api/auth/logout", controllers.Logout)
}

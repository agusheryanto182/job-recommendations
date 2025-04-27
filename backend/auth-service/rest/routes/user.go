package routes

import (
	"auth-service/internal/auth"
	"auth-service/internal/controllers"
	"auth-service/rest/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router, controller *controllers.AuthController) {
	// Google OAuth routes
	router.Get("/google/login", controller.GoogleLogin)
	router.Get("/google/callback", controller.GoogleCallback)

	router.Get("/profile", middlewares.Auth(auth.User), controller.GetUserProfile)
	router.Post("/refresh", middlewares.Auth(auth.User), controller.RefreshToken)
	router.Post("/logout", middlewares.Auth(auth.User), controller.Logout)
}

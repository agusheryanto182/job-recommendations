package routes

import (
	"auth-service/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router, controller *controllers.AuthController) {
	userRouter := router.Group("/user")

	// Google OAuth routes
	userRouter.Get("/google/login", controller.GoogleLogin)
	userRouter.Get("/google/callback", controller.GoogleCallback)

	userRouter.Get("/profile", controller.GetUserProfile)
}

package routes

import (
	"auth-service/internal/auth"
	"auth-service/internal/controllers"
	"auth-service/rest/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router, controller *controllers.AuthController) {
	authRouter := router.Group("/auth")
	userRouter := router.Group("/user")

	// Google OAuth routes
	authRouter.Get("/google/login", controller.GoogleLogin)
	authRouter.Get("/google/callback", controller.GoogleCallback)

	userRouter.Get("/profile", middlewares.Auth(auth.User), controller.GetUserProfile)
	userRouter.Post("/refresh", middlewares.Auth(auth.User), controller.RefreshToken)
	userRouter.Post("/logout", middlewares.Auth(auth.User), controller.Logout)
}

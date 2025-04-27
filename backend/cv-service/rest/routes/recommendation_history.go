package routes

import (
	"cv-service/internal/controllers"
	"cv-service/internal/grpc/client"
	"cv-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRecommendationHistoryRoutes(router fiber.Router, controller *controllers.RecommendationHistoryController, authClient *client.AuthClient) {
	router.Get("/user/history", middleware.AuthMiddleware(authClient), controller.FindByUserID)
}

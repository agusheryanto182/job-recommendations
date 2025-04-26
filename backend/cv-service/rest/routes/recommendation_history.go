package routes

import (
	"cv-service/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRecommendationHistoryRoutes(router fiber.Router, controller *controllers.RecommendationHistoryController) {
	recommendationHistoryRoutes := router.Group("/recommendation_history")

	recommendationHistoryRoutes.Get("/user/:user_id", controller.FindByUserID)
}

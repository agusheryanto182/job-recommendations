package controllers

import (
	"cv-service/internal/grpc/client"
	"cv-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type RecommendationHistoryController struct {
	recommendationHistoryService services.RecommendationHistoryService
	authClient                   *client.AuthClient
}

func NewRecommendationHistoryController(recommendationHistoryService services.RecommendationHistoryService, authClient *client.AuthClient) *RecommendationHistoryController {
	return &RecommendationHistoryController{
		recommendationHistoryService: recommendationHistoryService,
		authClient:                   authClient,
	}
}

func (c *RecommendationHistoryController) FindByUserID(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(string)

	recommendationHistory, err := c.recommendationHistoryService.FindByUserID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find recommendation_history",
		})
	}
	return ctx.JSON(fiber.Map{
		"recommendation_history": recommendationHistory,
	})
}

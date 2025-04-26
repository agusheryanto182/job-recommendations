package controllers

import (
	"cv-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type RecommendationHistoryController struct {
	recommendationHistoryService services.RecommendationHistoryService
}

func NewRecommendationHistoryController(recommendationHistoryService services.RecommendationHistoryService) *RecommendationHistoryController {
	return &RecommendationHistoryController{
		recommendationHistoryService: recommendationHistoryService,
	}
}

func (c *RecommendationHistoryController) FindByUserID(ctx *fiber.Ctx) error {
	userID := ctx.Params("user_id")
	recommendationHistory, err := c.recommendationHistoryService.FindByUserID(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"recommendation_history": recommendationHistory,
	})
}

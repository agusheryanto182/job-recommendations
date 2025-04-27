//go:build wireinject
// +build wireinject

package main

import (
	"cv-service/internal/controllers"
	"cv-service/internal/grpc/client"
	"cv-service/internal/repositories"
	"cv-service/internal/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var SuperSet = wire.NewSet(
	repositories.NewRecommendationHistoryRepo,
	services.NewRecommendationHistoryService,
	controllers.NewRecommendationHistoryController,
)

// Return pointer, not value
func InitializeRecommendationHistoryDependency(db *gorm.DB, authClient *client.AuthClient) *controllers.RecommendationHistoryController {
	wire.Build(SuperSet)
	return &controllers.RecommendationHistoryController{} // Return pointer
}

//go:build wireinject
// +build wireinject

package main

import (
	"auth-service/config"
	"auth-service/internal/controllers"
	"auth-service/internal/repositories"
	"auth-service/internal/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var SuperSet = wire.NewSet(
	repositories.NewUserRepo,
	services.NewAuthService,
	controllers.NewAuthController,
)

// Return pointer ke AuthController, bukan value
func InitializeUserDependency(db *gorm.DB, googleAuthCfg *config.GoogleOAuthConfig) *controllers.AuthController {
	wire.Build(SuperSet)
	return &controllers.AuthController{} // Return pointer
}

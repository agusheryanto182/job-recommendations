package controllers

import (
	"auth-service/internal/auth"
	"auth-service/internal/errs"
	"auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) GoogleCallback(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (c *AuthController) GoogleLogin(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (c *AuthController) GetUserProfile(ctx *fiber.Ctx) error {
	userID, err := auth.GetAuthID(ctx, auth.User)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	user, err := c.authService.GetUserProfile(userID)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return ctx.JSON(fiber.Map{
		"user": user,
	})
}

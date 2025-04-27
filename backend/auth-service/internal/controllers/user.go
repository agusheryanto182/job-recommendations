package controllers

import (
	"auth-service/config"
	"auth-service/internal/auth"
	"auth-service/internal/errs"
	"auth-service/internal/request"
	"auth-service/internal/services"
	"auth-service/pkg/uuid"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService   services.AuthService
	googleAuthCfg *config.GoogleOAuthConfig
}

func NewAuthController(authService services.AuthService, googleAuthCfg *config.GoogleOAuthConfig) *AuthController {
	return &AuthController{
		authService:   authService,
		googleAuthCfg: googleAuthCfg,
	}
}

func (c *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	oldToken, err := auth.GetJwtTokenFromHeader(ctx)
	if err != nil {
		return errs.NewUnauthorizedError(err.Error())
	}

	claims, err := auth.ParseJwt(oldToken, auth.User)
	if err != nil {
		return errs.NewUnauthorizedError(err.Error())
	}

	user, err := c.authService.GetUserProfile(claims.Id)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	newToken := auth.MakeJwt(user, auth.User)
	if err := auth.InvalidateToken(oldToken, auth.User); err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return ctx.JSON(fiber.Map{
		"message": "Token refreshed successfully",
		"token":   newToken,
		"user": fiber.Map{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"avatar": user.AvatarURL,
		},
	})
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	token, err := auth.GetJwtTokenFromHeader(ctx)
	if err != nil {
		return errs.NewUnauthorizedError(err.Error())
	}

	if err := auth.InvalidateToken(token, auth.User); err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return ctx.JSON(fiber.Map{
		"message": "Logout successful",
	})
}

func (c *AuthController) GoogleCallback(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	savedState := ctx.Cookies("oauth_state")
	if err := c.authService.ValidateGoogleState(state, savedState); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	code := ctx.Query("code")
	if code == "" {
		return errs.NewBadRequestError("Authorization code is required")
	}

	token, err := c.googleAuthCfg.GetConfig().Exchange(context.Background(), code)
	if err != nil {
		return errs.NewInternalServerError("Failed to exchange token")
	}

	userInfo, err := c.authService.GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		return err
	}

	userReq := request.AuthenticateUserRequest{
		ID:      userInfo.ID,
		Email:   userInfo.Email,
		Name:    userInfo.Name,
		Picture: userInfo.Picture,
	}

	user, err := c.authService.AuthenticateGoogle(userReq)
	if err != nil {
		return err
	}

	jwtToken := auth.MakeJwt(user, auth.User)

	return ctx.JSON(fiber.Map{
		"message": "Authentication successful",
		"token":   jwtToken,
		"user": fiber.Map{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"avatar": user.AvatarURL,
		},
	})

}

func (c *AuthController) GoogleLogin(ctx *fiber.Ctx) error {
	state := uuid.GenerateUUID()

	// save state in cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "lax",
	})

	// redirect to google
	url := c.googleAuthCfg.GetConfig().AuthCodeURL(state)

	return ctx.Redirect(url)
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

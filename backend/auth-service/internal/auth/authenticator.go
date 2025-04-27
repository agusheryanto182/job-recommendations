package auth

import (
	"auth-service/config"
	"auth-service/internal/errs"
	"auth-service/internal/models"
	"auth-service/internal/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Authenticable interface {
	GetName() string
	GetId() string
	GetEmail() string
}

func ValidateRequest(c *fiber.Ctx, guard Guard) (string, error) {
	token, err := GetJwtTokenFromHeader(c)
	if err != nil {
		return "", err
	}

	_, err = ParseJwt(token, guard)
	if err != nil {
		return "", err
	}

	err = ValidateToken(token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func InvalidateToken(token string, guard Guard) error {
	invalidTokenRepo, err := getInvalidTokenRepo()
	if err != nil {
		return err
	}

	claim, err := ParseJwt(token, guard)

	if err != nil {
		return err
	}

	_, err = invalidTokenRepo.Create(models.InvalidToken{
		Token:     token,
		ExpiredAt: time.Unix(claim.ExpiresAt, 0),
	})

	return err
}

func ValidateToken(token string) error {
	invalidTokenRepo, err := getInvalidTokenRepo()
	if err != nil {
		return err
	}

	_, err = invalidTokenRepo.FindByToken(token)

	// if token exist in table `invalid_token`
	if err == nil {
		return errs.NewUnauthorizedError("Token is invalid")
	}
	return nil
}

func getInvalidTokenRepo() (repositories.InvalidTokenRepository, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		// Load config error, send panic signal
		panic(errs.NewInternalServerError("Failed to load configuration"))
	}

	db, err := config.ConnectDB(cfg)
	if err != nil {
		// Connect DB error, send panic signal
		panic(errs.NewInternalServerError("Failed to connect database"))
	}

	invalidTokenRepo := repositories.NewInvalidTokenRepository(db)
	return invalidTokenRepo, nil
}

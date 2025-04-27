package middlewares

import (
	"auth-service/internal/auth"
	"auth-service/internal/errs"

	"github.com/gofiber/fiber/v2"
)

// Global auth middleware
func Auth(guard auth.Guard) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := auth.ValidateRequest(c, guard)
		if err != nil {
			return errs.NewUnauthorizedError(err.Error())
		}

		return c.Next()
	}
}

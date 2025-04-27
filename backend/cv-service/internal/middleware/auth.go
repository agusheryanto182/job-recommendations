package middleware

import (
	"cv-service/internal/errs"
	"cv-service/internal/grpc/client"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(authClient *client.AuthClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return errs.NewUnauthorizedError("Authorization header is missing")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token
		if err := authClient.ValidateToken(c.Context(), token); err != nil {
			return errs.NewUnauthorizedError(err.Error())
		}

		// Get user ID and store in context
		userID, err := authClient.GetUserID(c.Context(), token)
		if err != nil {
			return errs.NewUnauthorizedError(err.Error())
		}

		// Set user ID in context for later use
		c.Locals("userID", userID)

		return c.Next()
	}
}

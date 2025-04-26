package auth

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetJwtTokenFromHeader(c *fiber.Ctx) (string, error) {
	// Get Token
	token := c.Get("Authorization")

	if token == "" {
		return "", errors.New("no auth token found")
	}

	splitToken := strings.Split(token, "Bearer ")

	if len(splitToken) <= 1 {
		return "", errors.New("invalid token")
	}

	token = splitToken[1]

	return token, nil
}

func GetJwtClaimFromReq(c *fiber.Ctx, guard Guard) (*MyClaims, error) {
	// Get Token
	token, err := GetJwtTokenFromHeader(c)

	if err != nil {
		return nil, err
	}

	// Parse JWT
	return ParseJwt(token, guard)
}

func GetAuthID(c *fiber.Ctx, guard Guard) (string, error) {
	claim, err := GetJwtClaimFromReq(c, guard)
	if err != nil {
		return "", err
	}

	return claim.Id, nil
}

func GetAuthName(c *fiber.Ctx, guard Guard) (string, error) {
	claim, err := GetJwtClaimFromReq(c, guard)
	if err != nil {
		return "", err
	}

	return claim.Name, nil
}

func GetAuthEmail(c *fiber.Ctx, guard Guard) (string, error) {
	claim, err := GetJwtClaimFromReq(c, guard)
	if err != nil {
		return "", err
	}

	return claim.Email, nil
}

package auth

import (
	"auth-service/config"
	"auth-service/internal/errs"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Guard Guard  `json:"guard"`
	jwt.StandardClaims
}

func (m MyClaims) ValidateGuard(guard Guard) error {
	if m.Guard != guard {
		return errs.NewUnauthorizedError("Guard is invalid")
	}
	return nil
}

func ParseJwt(tokenString string, guard Guard) (*MyClaims, error) {
	token, err := parse(tokenString)

	if err != nil {
		return nil, errs.NewUnauthorizedError("jwt parse error")
	}

	claims, ok := token.Claims.(*MyClaims)

	if !ok {
		return nil, errs.NewUnauthorizedError("jwt parse error")
	}

	if err = claims.ValidateGuard(guard); err != nil {
		return nil, err
	}

	return claims, nil
}

func parse(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.GetPrivateKey()
	})
}

func MakeJwt(authenticable Authenticable, guard Guard) string {
	claims := makeMapClaims(authenticable, guard)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key, err := config.GetPrivateKey()

	if err != nil {
		// unexpected error found, send panic signal
		panic(errs.NewInternalServerError("Failed to get private key"))
	}

	ss, err := token.SignedString(key)

	if err != nil {
		panic(errs.NewInternalServerError("Failed to get signed string"))
	}

	return ss
}

func makeMapClaims(authenticalble Authenticable, guard Guard) jwt.Claims {
	var expiration time.Duration = guard.ExpireTime()

	return jwt.MapClaims{
		"id":    authenticalble.GetId(),
		"name":  authenticalble.GetName(),
		"email": authenticalble.GetEmail(),
		"guard": guard,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(expiration).Unix(),
	}
}

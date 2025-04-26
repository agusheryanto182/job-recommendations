package config

import (
	"errors"
	"os"
)

func GetPrivateKey() ([]byte, error) {
	secretStr := os.Getenv("JWT_SECRET_KEY")

	if secretStr == "" {
		return nil, errors.New("JWT_SECRET_KEY value in .env is EMPTY")
	}

	secretByte := []byte(secretStr)
	return secretByte, nil
}

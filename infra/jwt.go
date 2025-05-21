package infra

import (
	"fmt"
	"log"
	"msvc-users/application/contracts"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JWTEncrypter struct{}

func NewJWTEncrypter() contracts.Encrypter {
	return &JWTEncrypter{}
}

var _ contracts.Encrypter = (*JWTEncrypter)(nil)

func (e *JWTEncrypter) Encrypt(values interface{}) (string, error) {
	env := os.Getenv("ENV")
	if env == "" || env == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, using default env vars")
		}
	}

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	vals, ok := values.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid input format: expected map[string]interface{}")
	}

	claims := jwt.MapClaims{}
	for key, value := range vals {
		claims[key] = value
	}

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error generating JWT token: %w", err)
	}

	return tokenString, nil
}

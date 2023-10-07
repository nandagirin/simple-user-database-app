package utils

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(c *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")
	tokenHeader := c.GetReqHeaders()["Authorization"]
	tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return err
	}

	return nil
}

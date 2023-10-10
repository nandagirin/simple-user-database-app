package utils

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// This function will validate token from request header
// `Authorization` and check if the token is valid by
// comparing it with JWT_SECRET value. If the token is
// valid, it will return nil error. Otherwise, it will
// return error.
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

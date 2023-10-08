package handlers

import (
	"user/database"
	"user/models"
	utilities "user/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// UserList returns a list of users
func UserList(c *fiber.Ctx) error {
	err := utilities.ValidateToken(c)
	if err != nil {
		return Unauthorized(c, err.Error())
	}

	users := database.Get()

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	err := utilities.ValidateToken(c)
	if err != nil {
		return Unauthorized(c, err.Error())
	}

	user := &models.User{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("user")
		Name: utils.CopyString(c.FormValue("user")),
	}
	database.Insert(user)

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// NotFound returns custom 404 status code
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"message": "Not found",
	})
}

// Unauthorized returns custom 403 status code
func Unauthorized(c *fiber.Ctx, reason string) error {
	return c.Status(403).JSON(fiber.Map{
		"message": "Unauthorized",
		"reason":  reason,
	})
}

// Health check handler
func Health(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "UP",
	})
}

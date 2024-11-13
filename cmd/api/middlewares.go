package main

import (
	"docuSync/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func (app *Config) requestLogger(c *fiber.Ctx) error {
	start := time.Now()

	// Process request
	err := c.Next()

	app.logger.Log("info", "Request received", map[string]any{
		"Method":   c.Method(),
		"URL":      c.OriginalURL(),
		"Status":   c.Response().StatusCode(),
		"Duration": time.Since(start),
	})

	return err
}

func (app *Config) authenticate(c *fiber.Ctx) error {
	token := c.Get("authorization")

	if len(token) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "no token provided",
		})
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		log.Println(err)
		if errors.Is(err, utils.ErrInvalidToken) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "wrong token",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to authenticate user",
		})

	}

	c.Locals("user", userID)
	return c.Next()

}

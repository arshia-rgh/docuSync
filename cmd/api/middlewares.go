package main

import (
	"github.com/gofiber/fiber/v2"
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

}

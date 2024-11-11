package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserRegister struct {
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (app *Config) registerUser(c *fiber.Ctx) error {
	user := new(UserRegister)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data ",
			"error":   err.Error(),
		})
	}
}

package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

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
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data ",
			"error":   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	newUser, err := app.client.User.
		Create().
		SetName(user.Name).
		SetLastName(user.LastName).
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(newUser)

}

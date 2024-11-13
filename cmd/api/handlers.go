package main

import (
	"context"
	"docuSync/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

// registerUser uses the UserRegister schema
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

	hashPass, _ := utils.HashPassword(user.Password)
	newUser, err := app.client.User.
		Create().
		SetName(user.Name).
		SetLastName(user.LastName).
		SetUsername(user.Username).
		SetPassword(hashPass).
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

// loginUser uses the UserLogin schema
func (app *Config) loginUser(c *fiber.Ctx) error {
	var user UserLogin

	if err := c.BodyParser(user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": token,
	})
}

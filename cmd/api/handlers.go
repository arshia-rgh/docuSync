package main

import (
	"context"
	"docuSync/ent"
	UserDB "docuSync/ent/user"
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

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	dbUser, err := app.client.User.
		Query().
		Where(UserDB.UsernameEQ(user.Username)).
		Only(ctx)
	if err != nil {
		log.Println(err.Error())
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "failed to find user with given username",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})

	}

	ok := utils.VerifyPassword(user.Password, dbUser.Password)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong password",
		})
	}

	token, err := utils.GenerateToken(dbUser.ID)
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

func (app *Config) me(c *fiber.Ctx) error {
	userID := c.Locals("user").(int)

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	user, err := app.client.User.
		Get(ctx, userID)

	if err != nil {
		log.Println(err.Error())
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "no user found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})

	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (app *Config) updateUser(c *fiber.Ctx) error {
	var user UserUpdate
	if err := c.BodyParser(user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data ",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	dbUser, err := app.client.User.
		UpdateOneID(userID).
		SetName(user.Name).
		SetLastName(user.LastName).
		SetUsername(user.Username).
		SetEmail(user.Email).
		Save(ctx)

	if err != nil {
		log.Println(err.Error())
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "no user found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dbUser)
}

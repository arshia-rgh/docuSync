package main

import (
	"context"
	"docuSync/ent"
	DocumentDB "docuSync/ent/document"
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
		if ent.IsValidationError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "validation error",
				"error":   err.Error(),
			})
		}
		if ent.IsConstraintError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "username or email already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(newUser)

}

// loginUser uses the UserLogin schema
func (app *Config) loginUser(c *fiber.Ctx) error {
	user := new(UserLogin)
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

// me is protected by auth
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

// updateUser uses the UserUpdate schema and protected by auth
func (app *Config) updateUser(c *fiber.Ctx) error {
	user := new(UserUpdate)
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

	update := app.client.User.UpdateOneID(userID)
	if user.Name != "" {
		update.SetName(user.Name)
	}
	if user.LastName != "" {
		update.SetLastName(user.LastName)
	}
	if user.Username != "" {
		update.SetUsername(user.Username)
	}
	if user.Email != "" {
		update.SetEmail(user.Email)
	}

	dbUser, err := update.Save(ctx)
	if err != nil {
		log.Println(err.Error())
		if ent.IsValidationError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "validation error",
				"error":   err.Error(),
			})
		}
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

// changePassword uses ChangePassword schema and protected by auth
func (app *Config) changePassword(c *fiber.Ctx) error {
	newPasswordData := new(ChangePassword)

	if err := c.BodyParser(newPasswordData); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	user, err := app.client.User.Get(ctx, userID)

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

	if !utils.VerifyPassword(newPasswordData.OldPassword, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "your old password is wrong",
		})
	}
	if newPasswordData.NewPassword != newPasswordData.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "passwords arent same",
		})
	}

	newHashedPassword, err := utils.HashPassword(newPasswordData.NewPassword)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update password",
		})
	}

	dbUser, err := app.client.User.
		UpdateOneID(userID).
		SetPassword(newHashedPassword).
		Save(ctx)
	if err != nil {
		log.Println(err.Error())
		if ent.IsValidationError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "validation error",
				"error":   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "password changed successfully",
		"detail":  dbUser,
	})
}

// createDocument uses CreateDocument schema and protected by auth
func (app *Config) createDocument(c *fiber.Ctx) error {
	document := new(CreateDocument)

	if err := c.BodyParser(document); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}
	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	dbDocument, err := app.client.Document.
		Create().
		SetTitle(document.Title).
		SetOwnerID(userID).
		Save(ctx)

	if err != nil {
		log.Println(err.Error())
		if ent.IsConstraintError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "document with this title already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(dbDocument)
}

// changeDocumentTitle uses the ChangeDocumentTitle schema and protected by auth
func (app *Config) changeDocumentTitle(c *fiber.Ctx) error {
	document := new(ChangeDocumentTitle)

	if err := c.BodyParser(document); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}
	documentID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong id",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	dbDocument, err := app.client.Document.
		UpdateOneID(documentID).
		Where(DocumentDB.HasOwnerWith(UserDB.IDEQ(userID))).
		SetTitle(document.Title).
		Save(ctx)

	if err != nil {
		log.Println(err.Error())
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "no document found for the current user",
			})
		}
		if ent.IsConstraintError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "document with this title already exists",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(dbDocument)
}

// addDocumentText uses AddDocumentText schema and protected by auth
func (app *Config) addDocumentText(c *fiber.Ctx) error {
	document := new(AddDocumentText)

	if err := c.BodyParser(document); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}
	documentID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong id",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err = app.client.Document.
		UpdateOneID(documentID).
		Where(DocumentDB.Or(
			DocumentDB.HasOwnerWith(UserDB.IDEQ(userID)),
			DocumentDB.HasAllowedUsersWith(UserDB.IDEQ(userID)),
		)).
		SetText(document.Text).
		Save(ctx)

	if err != nil {
		log.Println(err.Error())
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "you are not owner of this document or dont have permission to this document",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	dbDocument, err := app.client.Document.
		UpdateOneID(documentID).
		AddEditorIDs(userID).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "document edited and saved successfully",
		"document": dbDocument,
	})
}

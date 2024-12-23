package main

import (
	"context"
	"docuSync/ent"
	DocumentDB "docuSync/ent/document"
	UserDB "docuSync/ent/user"
	"docuSync/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

const dbTimeout = time.Second * 3

// registerUser uses the UserRegister schema
func (cfg *Config) registerUser(c *fiber.Ctx) error {
	user := new(UserRegister)

	if err := c.BodyParser(user); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the registerUser handler failed", map[string]any{"err": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data ",
			"error":   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashPass, _ := utils.HashPassword(user.Password)
	newUser, err := cfg.Client.User.
		Create().
		SetName(user.Name).
		SetLastName(user.LastName).
		SetUsername(user.Username).
		SetPassword(hashPass).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		if ent.IsValidationError(err) {
			cfg.Logger.Log("warning", "Saving new user in the registerUser handler to the DB failed because of the validation error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "validation error",
				"error":   err.Error(),
			})
		}
		if ent.IsConstraintError(err) {
			cfg.Logger.Log("warning", "Saving new user in the registerUser handler to the DB failed because of the constraint error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "username or email already exists",
			})
		}
		cfg.Logger.Log("error", "Saving new user in the registerUser handler to the DB failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}
	cfg.Logger.Log("info", "New user successfully registered and added to the DB", map[string]any{"user": newUser})
	return c.Status(fiber.StatusOK).JSON(newUser)

}

// loginUser uses the UserLogin schema
func (cfg *Config) loginUser(c *fiber.Ctx) error {
	user := new(UserLogin)
	if err := c.BodyParser(user); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the loginUser handler failed", map[string]any{"err": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	dbUser, err := cfg.Client.User.
		Query().
		Where(UserDB.UsernameEQ(user.Username)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			cfg.Logger.Log("warning", "Retrieving user in the loginUser handler from the DB failed because of the NotFound error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "failed to find user with given username",
			})
		}
		cfg.Logger.Log("error", "Retrieving user in the loginUser handler from the DB failed because of the server error", map[string]any{"error": err.Error()})
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
		cfg.Logger.Log("error", "generate token for the user in the loginUser handler failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	cfg.Logger.Log("info", "New token generated and sent to the user", map[string]any{"userID": dbUser.ID, "token": token})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": token,
	})
}

// me is protected by auth
func (cfg *Config) me(c *fiber.Ctx) error {
	userID := c.Locals("user").(int)

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	user, err := cfg.Client.User.
		Get(ctx, userID)

	if err != nil {
		cfg.Logger.Log("warning", "Retrieving user in the me handler from the DB failed because of the NotFound error", map[string]any{"error": err.Error()})
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "no user found",
			})
		}
		cfg.Logger.Log("error", "Retrieving user in the me handler from the DB failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})

	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// updateUser uses the UserUpdate schema and protected by auth
func (cfg *Config) updateUser(c *fiber.Ctx) error {
	user := new(UserUpdate)
	if err := c.BodyParser(user); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the updateUser handler failed", map[string]any{"err": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data ",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	update := cfg.Client.User.UpdateOneID(userID)
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
		if ent.IsValidationError(err) {
			cfg.Logger.Log("warning", "updating user in the updateUser handler to the DB failed because of the validation error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "validation error",
				"error":   err.Error(),
			})
		}
		if ent.IsNotFound(err) {
			cfg.Logger.Log("warning", "updating user in the updateUser handler to the DB failed because of the NotFound error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "no user found",
			})
		}
		cfg.Logger.Log("error", "updating user in the updateUser handler to the DB failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	cfg.Logger.Log("info", "User in the updateUser handler successfully updated", map[string]any{"userID": userID})
	return c.Status(fiber.StatusOK).JSON(dbUser)
}

// changePassword uses ChangePassword schema and protected by auth
func (cfg *Config) changePassword(c *fiber.Ctx) error {
	newPasswordData := new(ChangePassword)

	if err := c.BodyParser(newPasswordData); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the changePassword handler failed", map[string]any{"err": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	user, err := cfg.Client.User.Get(ctx, userID)

	if err != nil {
		if ent.IsNotFound(err) {
			cfg.Logger.Log("warning", "Retrieving user in the changePassword handler failed because of the NotFound error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "no user found",
			})
		}
		cfg.Logger.Log("error", "Retrieving user in the changePassword handler failed because of the server error", map[string]any{"error": err.Error()})
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
		cfg.Logger.Log("error", "Hashing password failed in the changePassword handler", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update password",
		})
	}

	dbUser, err := cfg.Client.User.
		UpdateOneID(userID).
		SetPassword(newHashedPassword).
		Save(ctx)
	if err != nil {
		if ent.IsValidationError(err) {
			cfg.Logger.Log("warning", "Changing password of the user in the changePassword handler failed because of the validation error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "validation error",
				"error":   err.Error(),
			})
		}
		cfg.Logger.Log("error", "Changing password of the user in the changePassword handler failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	cfg.Logger.Log("error", "Changing password of the user in the changePassword handler successfully completed", map[string]any{"userID": userID})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "password changed successfully",
		"detail":  dbUser,
	})
}

// createDocument uses CreateDocument schema and protected by auth
func (cfg *Config) createDocument(c *fiber.Ctx) error {
	document := new(CreateDocument)

	if err := c.BodyParser(document); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the createDocument handler failed", map[string]any{"err": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
			"error":   err.Error(),
		})
	}
	userID := c.Locals("user").(int)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	dbDocument, err := cfg.Client.Document.
		Create().
		SetTitle(document.Title).
		SetOwnerID(userID).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			cfg.Logger.Log("warning", "Saving new document in the createDocument handler to the DB failed because of the constraint error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "document with this title already exists",
			})
		}

		cfg.Logger.Log("error", "Saving new document in the createDocument handler to the DB failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	cfg.Logger.Log("info", "Saving new document in the createDocument handler to the DB successfully completed", map[string]any{"document": dbDocument})
	return c.Status(fiber.StatusOK).JSON(dbDocument)
}

// changeDocumentTitle uses the ChangeDocumentTitle schema and protected by auth
func (cfg *Config) changeDocumentTitle(c *fiber.Ctx) error {
	document := new(ChangeDocumentTitle)

	if err := c.BodyParser(document); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the changeDocumentTitle handler failed", map[string]any{"err": err.Error()})
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

	dbDocument, err := cfg.Client.Document.
		UpdateOneID(documentID).
		Where(DocumentDB.HasOwnerWith(UserDB.IDEQ(userID))).
		SetTitle(document.Title).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			cfg.Logger.Log("warning", "Changing document title in the changeDocumentTitle handler failed because of the NotFound error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "no document found for the current user",
			})
		}
		if ent.IsConstraintError(err) {
			cfg.Logger.Log("warning", "Changing document title in the changeDocumentTitle handler failed because of the constrain error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "document with this title already exists",
			})
		}

		cfg.Logger.Log("error", "Changing document title in the changeDocumentTitle handler failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	cfg.Logger.Log("error", "Changing document title in the changeDocumentTitle handler successfully completed", map[string]any{"document": dbDocument})
	return c.Status(fiber.StatusOK).JSON(dbDocument)
}

// addDocumentText uses AddDocumentText schema and protected by auth
func (cfg *Config) addDocumentText(c *fiber.Ctx) error {
	document := new(AddDocumentText)

	if err := c.BodyParser(document); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the addDocumentText handler failed", map[string]any{"err": err.Error()})
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

	_, err = cfg.Client.Document.
		UpdateOneID(documentID).
		Where(DocumentDB.Or(
			DocumentDB.HasOwnerWith(UserDB.IDEQ(userID)),
			DocumentDB.HasAllowedUsersWith(UserDB.IDEQ(userID)),
		)).
		SetText(document.Text).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			cfg.Logger.Log("warning", "adding text to the document in the addDocumentText handler failed because of the NotFound error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "you are not owner of this document or dont have permission to this document",
			})
		}
		cfg.Logger.Log("error", "adding text to the document  in the addDocumentText handler failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	dbDocument, err := cfg.Client.Document.
		UpdateOneID(documentID).
		AddEditorIDs(userID).
		Save(ctx)
	if err != nil {
		cfg.Logger.Log("error", "adding new editor to the document in the addDocumentText handler failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	cfg.Logger.Log("info", "adding text to the document in the addDocumentText handler successfully completed", map[string]any{"document": dbDocument})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "document edited and saved successfully",
		"document": dbDocument,
	})
}

// addUserToTheAllowedEditorsOfDocument uses AddUserToTheAllowedEditorsOfDocument schema and protected by auth
func (cfg *Config) addUserToTheAllowedEditorsOfDocument(c *fiber.Ctx) error {
	data := new(AddUserToTheAllowedEditorsOfDocument)

	if err := c.BodyParser(data); err != nil {
		cfg.Logger.Log("warning", "Parsing data in the addUserToTheAllowedEditorsOfDocument handler failed", map[string]any{"err": err.Error()})
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

	dbDocument, err := cfg.Client.Document.
		UpdateOneID(documentID).
		Where(DocumentDB.HasOwnerWith(UserDB.IDEQ(userID))).
		AddAllowedUserIDs(data.UserID).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			cfg.Logger.Log("warning", "adding user as allowed editor to the document in the addUserToTheAllowedEditorsOfDocument handler failed because of the NotFound error", map[string]any{"error": err.Error()})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "you are not owner of this document",
			})
		}
		cfg.Logger.Log("error", "adding user as allowed editor to the document in the addUserToTheAllowedEditorsOfDocument handler failed because of the server error", map[string]any{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	cfg.Logger.Log("info", "adding user as allowed editor to the document in the addUserToTheAllowedEditorsOfDocument handler successfully completed", map[string]any{"documentID": documentID, "userID": userID})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("the user with %v id added as allowed editor to the document with %v id", userID, dbDocument.ID),
	})

}

package main

import (
	"context"
	"docuSync/ent/enttest"
	"docuSync/utils"
	"github.com/gofiber/fiber/v2"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestApp(t *testing.T) (*fiber.App, *Config) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	t.Cleanup(func() {
		client.Close()
	})

	server := fiber.New()
	app := &Config{
		client: client,
	}

	return server, app
}

func setupTestUser(t *testing.T) {
	_, app := setupTestApp(t)

	hashedPass, _ := utils.HashPassword("password123")
	_, _ = app.client.User.Create().
		SetName("John").
		SetLastName("Doe").
		SetUsername("johndoe").
		SetPassword(hashedPass).
		SetEmail("johndoe@example.com").Save(context.Background())
}

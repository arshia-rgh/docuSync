package main

import (
	"docuSync/ent/enttest"
	"github.com/gofiber/fiber/v2"
	"testing"
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

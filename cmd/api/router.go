package main

import "github.com/gofiber/fiber/v2"

func (app *Config) registerRouter(server *fiber.App) {
	server.Post("/register", app.registerUser)
}

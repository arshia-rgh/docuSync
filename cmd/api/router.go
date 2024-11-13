package main

import "github.com/gofiber/fiber/v2"

func (app *Config) registerPublicRouter(server fiber.Router) {
	server.Post("/register", app.registerUser)
}

func (app *Config) registerProtectedRouter(server fiber.Router) {
}

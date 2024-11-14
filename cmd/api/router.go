package main

import "github.com/gofiber/fiber/v2"

func (app *Config) registerPublicRouter(server fiber.Router) {
	server.Post("/register", app.registerUser)
	server.Post("/login", app.loginUser)
}

func (app *Config) registerProtectedRouter(server fiber.Router) {
	server.Get("/me", app.me)
	server.Put("/me/update", app.updateUser)
	server.Post("/me/change-password", app.changePassword)
	// document handlers
	server.Post("/create-document", app.createDocument)
}

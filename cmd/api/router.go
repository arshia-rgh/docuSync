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
	server.Post("/document/create", app.createDocument)
	server.Put("/document/change-title/:id", app.changeDocumentTitle)
	server.Post("/document/add-text/:id", app.addDocumentText)
	server.Post("/document/add-allowed-user/:id", app.addUserToTheAllowedEditorsOfDocument)
}

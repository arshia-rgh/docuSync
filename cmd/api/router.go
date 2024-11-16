package main

import (
	"github.com/gofiber/fiber/v2"
)

func (cfg *Config) registerPublicRouter(server fiber.Router) {
	server.Post("/register", cfg.registerUser)
	server.Post("/login", cfg.loginUser)
}

func (cfg *Config) registerProtectedRouter(server fiber.Router) {
	server.Get("/me", cfg.me)
	server.Put("/me/update", cfg.updateUser)
	server.Post("/me/change-password", cfg.changePassword)
	// document handlers
	server.Post("/document/create", cfg.createDocument)
	server.Put("/document/change-title/:id", cfg.changeDocumentTitle)
	server.Post("/document/add-text/:id", cfg.addDocumentText)
	server.Post("/document/add-user/:id", cfg.addUserToTheAllowedEditorsOfDocument)
}

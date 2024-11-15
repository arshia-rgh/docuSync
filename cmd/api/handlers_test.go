package main

import (
	"bytes"
	"context"
	"docuSync/ent"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	server, app := setupTestApp(t)
	server.Post("/register", app.registerUser)

	user := UserRegister{
		Name:     "John",
		LastName: "Doe",
		Username: "johndoe",
		Password: "password123",
		Email:    "johndoe@example.com",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := server.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseUser ent.User
	bodyBytes, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(bodyBytes, &responseUser)

	dbUser, err := app.client.User.Get(context.Background(), responseUser.ID)
	if err != nil {
		t.Fatalf("Failed to get user from database: %v", err)
	}
	assert.Equal(t, user.Username, dbUser.Username)
	assert.Equal(t, user.Email, dbUser.Email)

}

func TestLoginUser(t *testing.T) {
	server, app := setupTestApp(t)
	server.Post("/register", app.registerUser)
	server.Post("/login", app.loginUser)

	// Register the user first
	user := UserRegister{
		Name:     "John",
		LastName: "Doe",
		Username: "johndoe",
		Password: "password123",
		Email:    "johndoe@example.com",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := server.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Now login with the registered user
	loginData := UserLogin{
		Username: "johndoe",
		Password: "password123",
	}
	body, _ = json.Marshal(loginData)
	req = httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err = server.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseLogin map[string]string
	bodyBytes, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(bodyBytes, &responseLogin)

	_, exists := responseLogin["code"]
	assert.True(t, exists, "The `code` key should exists in the response")
}

func TestUpdateUser(t *testing.T) {
	server, app := setupTestApp(t)
	protectedServer := server.Group("/protected")
	protectedServer.Use(app.authenticate)
	protectedServer.Put("/me/update", app.updateUser)
	server.Post("/login", app.loginUser)

	setupTestUser(t)

	loginData := UserLogin{
		Username: "johndoe",
		Password: "password123",
	}
	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := server.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}
	var response map[string]string
	bodyBytes, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(bodyBytes, &response)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	code := response["code"]

	updateUserData := UserUpdate{
		Name: "test John",
	}
	body, _ = json.Marshal(updateUserData)
	req = httptest.NewRequest("PUT", "/protected/me/update", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", code)
	resp, err = server.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}

	var responseUser ent.User
	bodyBytes, _ = io.ReadAll(resp.Body)
	_ = json.Unmarshal(bodyBytes, &responseUser)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, updateUserData.Name, responseUser.Name)

	dbUser, err := app.client.User.Get(context.Background(), responseUser.ID)
	if err != nil {
		t.Fatalf("Failed to get user from database: %v", err)
	}

	assert.Equal(t, updateUserData.Name, dbUser.Name)
}

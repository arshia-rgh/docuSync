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

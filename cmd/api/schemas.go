package main

import "time"

type UserRegister struct {
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type ChangePassword struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreateDocument struct {
	Title string `json:"title,omitempty"`
}

type ChangeDocumentTitle struct {
	Title string `json:"title"`
}

type AddDocumentText struct {
	Text string `json:"text"`
}

type AddUserToTheAllowedEditorsOfDocument struct {
	UserID int `json:"user_id"`
}

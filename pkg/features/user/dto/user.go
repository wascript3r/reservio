package dto

import (
	"html"
)

// Create

type CreateReq struct {
	Email    string `json:"email" validate:"required,u_email"`
	Password string `json:"password" validate:"required,u_password"`
}

func (c *CreateReq) Escape() {
	c.Email = html.EscapeString(c.Email)
}

// Authenticate

type AuthenticateReq struct {
	Email    string `json:"email" validate:"required,u_email"`
	Password string `json:"password" validate:"required,u_password"`
}

type AuthenticateRes struct {
	Token string `json:"token"`
}

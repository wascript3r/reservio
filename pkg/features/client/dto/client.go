package dto

import (
	"html"

	udto "github.com/wascript3r/reservio/pkg/features/user/dto"
)

// Create

type CreateReq struct {
	udto.CreateReq
	FirstName string `json:"firstName" validate:"required,c_first_name"`
	LastName  string `json:"lastName" validate:"required,c_last_name"`
	Phone     string `json:"phone" validate:"required,c_phone"`
}

func (c *CreateReq) Escape(escapeUser bool) {
	if escapeUser {
		c.CreateReq.Escape()
	}
	c.FirstName = html.EscapeString(c.FirstName)
	c.LastName = html.EscapeString(c.LastName)
	c.Phone = html.EscapeString(c.Phone)
}

type Client struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type CreateRes Client

// Get

type ClientReq struct {
	ClientID string `json:"-" validate:"required,uuid"`
}

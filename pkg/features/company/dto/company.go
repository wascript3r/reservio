package dto

import (
	"html"

	udto "github.com/wascript3r/reservio/pkg/features/user/dto"
)

// Create

type CreateReq struct {
	udto.CreateReq
	Name        string `json:"name" validate:"required,c_name"`
	Address     string `json:"address" validate:"required,c_address"`
	Description string `json:"description" validate:"required,c_description"`
}

func (c *CreateReq) Escape(escapeUser bool) {
	if escapeUser {
		c.CreateReq.Escape()
	}
	c.Name = html.EscapeString(c.Name)
	c.Address = html.EscapeString(c.Address)
	c.Description = html.EscapeString(c.Description)
}

type Company struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Approved    bool   `json:"approved"`
}

type CreateRes Company

// Get

type CompanyReq struct {
	CompanyID string `json:"-" validate:"required,uuid"`
}

type GetReq CompanyReq

type GetRes Company

// GetAll

type GetAllRes struct {
	Companies []*Company `json:"companies"`
}

// Update

type UpdateReq struct {
	GetReq
	Name        *string `json:"name" validate:"omitempty,c_name"`
	Address     *string `json:"address" validate:"omitempty,c_address"`
	Description *string `json:"description" validate:"omitempty,c_description"`
	Approved    *bool   `json:"approved" validate:"omitempty"`
}

func (u *UpdateReq) Escape() {
	if u.Name != nil {
		*u.Name = html.EscapeString(*u.Name)
	}
	if u.Address != nil {
		*u.Address = html.EscapeString(*u.Address)
	}
	if u.Description != nil {
		*u.Description = html.EscapeString(*u.Description)
	}
}

// Delete

type DeleteReq GetReq

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
	*AdminUpdate
	*CompanyUpdate
}

type AdminUpdate struct {
	Approved *bool `json:"approved" validate:"omitempty"`
}

type CompanyUpdate struct {
	Name        *string `json:"name" validate:"omitempty,c_name"`
	Address     *string `json:"address" validate:"omitempty,c_address"`
	Description *string `json:"description" validate:"omitempty,c_description"`
}

func (u *UpdateReq) Escape() {
	if u.CompanyUpdate != nil {
		if u.CompanyUpdate.Name != nil {
			*u.CompanyUpdate.Name = html.EscapeString(*u.CompanyUpdate.Name)
		}
		if u.CompanyUpdate.Address != nil {
			*u.CompanyUpdate.Address = html.EscapeString(*u.CompanyUpdate.Address)
		}
		if u.CompanyUpdate.Description != nil {
			*u.CompanyUpdate.Description = html.EscapeString(*u.CompanyUpdate.Description)
		}
	}
}

// Delete

type DeleteReq GetReq

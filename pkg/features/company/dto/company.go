package dto

import (
	udto "github.com/wascript3r/reservio/pkg/features/user/dto"
)

// Create

type CreateReq struct {
	udto.CreateReq
	Name        string `json:"name" validate:"required,c_name"`
	Address     string `json:"address" validate:"required,c_address"`
	Description string `json:"description" validate:"required,c_description"`
}

type CreateRes struct {
	ID string `json:"id"`
}

// GetAll

type Company struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type GetAllRes struct {
	Companies []*Company `json:"companies"`
}

// Get

type GetReq struct {
	CompanyID string `json:"-" validate:"required,uuid"`
}

type GetRes Company

// Update

type UpdateReq struct {
	GetReq
	Name        *string `json:"name" validate:"omitempty,c_name"`
	Address     *string `json:"address" validate:"omitempty,c_address"`
	Description *string `json:"description" validate:"omitempty,c_description"`
}

// Delete

type DeleteReq GetReq

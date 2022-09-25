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
	CompanyID string `json:"companyID"`
}

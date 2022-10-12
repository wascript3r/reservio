package dto

import (
	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

type AccessClaims struct {
	UserID         string       `json:"userID"`
	Role           umodels.Role `json:"role"`
	RefreshTokenID string       `json:"rtID"`
}

type RefreshClaims struct {
	RefreshTokenID string `json:"rtID"`
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// IssueAccess

type RenewAccessReq struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RenewAccessRes struct {
	AccessToken string `json:"accessToken"`
}

package token

import (
	"github.com/wascript3r/reservio/pkg/features/user/models"
)

type UserClaims struct {
	UserID string `json:"userID"`
	Role   string `json:"role"`
}

type Usecase interface {
	Generate(us *models.User) (token string, err error)
	Parse(token string) (*UserClaims, error)
}

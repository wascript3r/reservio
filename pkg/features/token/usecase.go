package token

import (
	"context"

	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

type UserClaims struct {
	UserID string       `json:"userID"`
	Role   umodels.Role `json:"role"`
}

type Usecase interface {
	Generate(us *umodels.User) (token string, err error)
	Parse(token string) (*UserClaims, error)
	StoreCtx(ctx context.Context, claims *UserClaims) context.Context
	LoadCtx(ctx context.Context) (*UserClaims, error)
}

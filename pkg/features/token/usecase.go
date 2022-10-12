package token

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/token/dto"
	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

type Usecase interface {
	IssuePair(ctx context.Context, us *umodels.User) (*dto.TokenPair, error)
	RenewAccess(ctx context.Context, req *dto.RenewAccessReq) (*dto.RenewAccessRes, error)
	ParseAccess(tkn string) (*dto.AccessClaims, error)
	ParseRefresh(tkn string) (*dto.RefreshClaims, error)
	StoreCtx(ctx context.Context, claims *dto.AccessClaims) context.Context
	LoadCtx(ctx context.Context) (*dto.AccessClaims, error)
}

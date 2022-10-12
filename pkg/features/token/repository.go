package token

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/token/models"
)

type Repository interface {
	Insert(ctx context.Context, rts *models.RefreshToken) (id string, err error)
	GetClaims(ctx context.Context, refreshTokenID string) (*models.Claims, error)
	Delete(ctx context.Context, refreshTokenID string) error
}

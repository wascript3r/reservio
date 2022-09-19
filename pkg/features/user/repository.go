package user

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/user/models"
)

type Repository interface {
	Insert(ctx context.Context, us *models.User) (id string, err error)
	EmailExists(ctx context.Context, email string) (bool, error)
}

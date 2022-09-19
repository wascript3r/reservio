package user

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/company/models"
)

type Repository interface {
	Insert(ctx context.Context, cs *models.Company) error
}

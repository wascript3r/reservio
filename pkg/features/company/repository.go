package company

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/company/models"
)

type Repository interface {
	Insert(ctx context.Context, cs *models.Company) error
	Get(ctx context.Context, userID string) (*models.CompanyInfo, error)
	GetAll(ctx context.Context) ([]*models.CompanyInfo, error)
	Update(ctx context.Context, userID string, cu *models.CompanyUpdate) error
	Delete(ctx context.Context, userID string) error
}

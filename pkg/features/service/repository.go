package service

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/service/models"
)

type Repository interface {
	Insert(ctx context.Context, ss *models.Service) (id string, err error)
	Get(ctx context.Context, companyID, serviceID string, onlyApprovedCompany bool) (*models.Service, error)
	GetAll(ctx context.Context, companyID string, onlyApprovedCompany bool) ([]*models.Service, error)
	Update(ctx context.Context, companyID, serviceID string, su *models.ServiceUpdate) error
}

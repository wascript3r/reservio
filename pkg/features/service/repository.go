package service

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/service/models"
)

type Repository interface {
	Insert(ctx context.Context, ss *models.Service) (id string, err error)
	Get(ctx context.Context, companyID, serviceID string, onlyApprovedCompany bool) (*models.Service, error)
}

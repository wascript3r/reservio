package client

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/client/models"
)

type Repository interface {
	Insert(ctx context.Context, cs *models.Client) error
	Get(ctx context.Context, id string) (*models.Client, error)
}

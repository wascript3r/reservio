package client

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/client/dto"
)

type Usecase interface {
	Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error)
}

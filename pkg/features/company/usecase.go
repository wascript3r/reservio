package company

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/company/dto"
)

type Usecase interface {
	Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error)
	Get(ctx context.Context, req *dto.GetReq, onlyApproved bool) (*dto.GetRes, error)
	GetAll(ctx context.Context, onlyApproved bool) (*dto.GetAllRes, error)
	Update(ctx context.Context, req *dto.UpdateReq) error
	Delete(ctx context.Context, req *dto.DeleteReq) error
}

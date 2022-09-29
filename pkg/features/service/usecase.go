package service

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/service/dto"
)

type Usecase interface {
	Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error)
	Get(ctx context.Context, req *dto.GetReq, onlyApprovedCompany bool) (*dto.GetRes, error)
	GetAll(ctx context.Context, req *dto.GetAllReq, onlyApprovedCompany bool) (*dto.GetAllRes, error)
	Update(ctx context.Context, req *dto.UpdateReq) error
}

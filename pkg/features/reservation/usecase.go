package reservation

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/reservation/dto"
)

type Usecase interface {
	Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error)
	Get(ctx context.Context, req *dto.GetReq, onlyApprovedCompany bool) (*dto.GetRes, error)
	GetAll(ctx context.Context, req *dto.GetAllReq, onlyApprovedCompany bool) (*dto.GetAllRes, error)
	GetAllMeta(ctx context.Context, req *dto.GetAllMetaReq, onlyApprovedCompany bool) (*dto.GetAllMetaRes, error)
	GetAllByClient(ctx context.Context, req *dto.GetAllByClientReq) (*dto.GetAllByClientRes, error)
	Update(ctx context.Context, req *dto.UpdateReq) error
	Delete(ctx context.Context, req *dto.DeleteReq) error
}

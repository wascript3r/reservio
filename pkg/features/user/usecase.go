package user

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/user/dto"
	"github.com/wascript3r/reservio/pkg/features/user/models"
)

type Usecase interface {
	Create(ctx context.Context, req *dto.CreateReq, role models.Role, validateReq bool) (id string, err error)
}

package user

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/user/dto"
)

type Usecase interface {
	Create(ctx context.Context, req *dto.CreateReq, validateReq bool) (id string, err error)
}

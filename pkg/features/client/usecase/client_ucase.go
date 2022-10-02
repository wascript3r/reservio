package usecase

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/client"
	"github.com/wascript3r/reservio/pkg/features/client/dto"
	"github.com/wascript3r/reservio/pkg/features/client/models"
	"github.com/wascript3r/reservio/pkg/features/user"
	"github.com/wascript3r/reservio/pkg/repository"
)

type Usecase struct {
	tx         repository.Transactor
	clientRepo client.Repository
	ctxTimeout time.Duration

	validator client.Validator
	userUcase user.Usecase
}

func New(tx repository.Transactor, cr client.Repository, t time.Duration, v client.Validator, uu user.Usecase) *Usecase {
	return &Usecase{
		tx:         tx,
		clientRepo: cr,
		ctxTimeout: t,

		validator: v,
		userUcase: uu,
	}
}

func (u *Usecase) Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, client.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var clientID string
	err := u.tx.WithinTx(c, func(c context.Context) error {
		var err error
		clientID, err = u.userUcase.Create(c, &req.CreateReq, false)
		if err != nil {
			return err
		}

		req.Escape(false)
		cs := &models.Client{
			ID:        clientID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Phone:     req.Phone,
		}

		return u.clientRepo.Insert(c, cs)
	})
	if err != nil {
		return nil, err
	}

	return &dto.CreateRes{
		ID: clientID,
	}, nil
}

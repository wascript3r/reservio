package usecase

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/repository"

	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/features/service/dto"
	"github.com/wascript3r/reservio/pkg/features/service/models"
)

type Usecase struct {
	serviceRepo service.Repository
	ctxTimeout  time.Duration

	validator service.Validator
}

func New(sr service.Repository, t time.Duration, v service.Validator) *Usecase {
	return &Usecase{
		serviceRepo: sr,
		ctxTimeout:  t,

		validator: v,
	}
}

func (u *Usecase) Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, service.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ws := make(models.WorkSchedule)
	for k, v := range req.WorkSchedule {
		ws[k] = &models.WorkHours{
			From: v.From,
			To:   v.To,
		}
	}

	ss := &models.Service{
		CompanyID:       req.CompanyID,
		Title:           req.Title,
		Description:     req.Description,
		SpecialistName:  req.SpecialistName,
		SpecialistPhone: req.SpecialistPhone,
		WorkSchedule:    ws,
	}

	id, err := u.serviceRepo.Insert(c, ss)
	if err != nil {
		if err == repository.ErrNoRelatedItems {
			return nil, company.NotFoundError
		}
		return nil, err
	}

	return &dto.CreateRes{
		ID: id,
	}, nil
}

package usecase

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/features/service/dto"
	"github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/repository"
)

type Usecase struct {
	serviceRepo service.Repository
	companyRepo company.Repository
	ctxTimeout  time.Duration

	validator service.Validator
}

func New(sr service.Repository, cr company.Repository, t time.Duration, v service.Validator) *Usecase {
	return &Usecase{
		serviceRepo: sr,
		companyRepo: cr,
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

func (u *Usecase) Get(ctx context.Context, req *dto.GetReq, onlyApprovedCompany bool) (*dto.GetRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, service.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	_, err := u.companyRepo.Get(c, req.CompanyID, onlyApprovedCompany)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, company.NotFoundError
		}
		return nil, err
	}

	ss, err := u.serviceRepo.Get(c, req.CompanyID, req.ServiceID, onlyApprovedCompany)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, service.NotFoundError
		}
		return nil, err
	}

	ws := make(dto.WorkSchedule)
	for k, v := range ss.WorkSchedule {
		ws[k] = (*dto.WorkHours)(v)
	}

	return &dto.GetRes{
		ID:              ss.ID,
		CompanyID:       ss.CompanyID,
		Title:           ss.Title,
		Description:     ss.Description,
		SpecialistName:  ss.SpecialistName,
		SpecialistPhone: ss.SpecialistPhone,
		WorkSchedule:    ws,
	}, nil
}

func (u *Usecase) GetAll(ctx context.Context, req *dto.GetAllReq, onlyApprovedCompany bool) (*dto.GetAllRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, service.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	_, err := u.companyRepo.Get(c, req.CompanyID, onlyApprovedCompany)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, company.NotFoundError
		}
		return nil, err
	}

	ss, err := u.serviceRepo.GetAll(c, req.CompanyID, onlyApprovedCompany)
	if err != nil {
		return nil, err
	}

	res := &dto.GetAllRes{
		Services: make([]*dto.Service, 0, len(ss)),
	}
	for _, s := range ss {
		ws := make(dto.WorkSchedule)
		for k, v := range s.WorkSchedule {
			ws[k] = (*dto.WorkHours)(v)
		}

		res.Services = append(res.Services, &dto.Service{
			ID:              s.ID,
			CompanyID:       s.CompanyID,
			Title:           s.Title,
			Description:     s.Description,
			SpecialistName:  s.SpecialistName,
			SpecialistPhone: s.SpecialistPhone,
			WorkSchedule:    ws,
		})
	}

	return res, nil
}
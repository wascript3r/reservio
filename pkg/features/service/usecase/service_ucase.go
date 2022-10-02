package usecase

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/features/service/dto"
	"github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/repository"
)

type Usecase struct {
	tx              repository.Transactor
	serviceRepo     service.Repository
	reservationRepo reservation.Repository
	companyRepo     company.Repository
	ctxTimeout      time.Duration

	validator service.Validator
}

func New(tx repository.Transactor, sr service.Repository, rs reservation.Repository, cr company.Repository, t time.Duration, v service.Validator) *Usecase {
	return &Usecase{
		tx:              tx,
		serviceRepo:     sr,
		reservationRepo: rs,
		companyRepo:     cr,
		ctxTimeout:      t,

		validator: v,
	}
}

func (u *Usecase) Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, service.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	_, err := u.companyRepo.Get(c, req.CompanyID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, company.NotFoundError
		}
		return nil, err
	}

	req.Escape()
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
		VisitDuration:   req.VisitDuration,
		WorkSchedule:    ws,
	}

	id, err := u.serviceRepo.Insert(c, ss)
	if err != nil {
		if err == repository.ErrConflictWithRelatedItems {
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
		VisitDuration:   ss.VisitDuration,
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
		Services: make([]*dto.Service, len(ss)),
	}
	for i, s := range ss {
		ws := make(dto.WorkSchedule)
		for k, v := range s.WorkSchedule {
			ws[k] = (*dto.WorkHours)(v)
		}

		res.Services[i] = &dto.Service{
			ID:              s.ID,
			CompanyID:       s.CompanyID,
			Title:           s.Title,
			Description:     s.Description,
			SpecialistName:  s.SpecialistName,
			SpecialistPhone: s.SpecialistPhone,
			VisitDuration:   s.VisitDuration,
			WorkSchedule:    ws,
		}
	}

	return res, nil
}

func (u *Usecase) Update(ctx context.Context, req *dto.UpdateReq) error {
	if err := u.validator.RawRequest(req); err != nil {
		return service.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	_, err := u.companyRepo.Get(c, req.CompanyID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return company.NotFoundError
		}
		return err
	}

	req.Escape()
	var ws *models.WorkSchedule
	if req.WorkSchedule != nil {
		w := make(models.WorkSchedule)
		ws = &w
		for k, v := range *req.WorkSchedule {
			w[k] = &models.WorkHours{
				From: v.From,
				To:   v.To,
			}
		}
	}

	su := &models.ServiceUpdate{
		Title:        req.Title,
		Description:  req.Description,
		WorkSchedule: ws,
	}
	if req.SpecialistName != nil {
		su.SpecialistName = &req.SpecialistName.Value
	}
	if req.SpecialistPhone != nil {
		su.SpecialistPhone = &req.SpecialistPhone.Value
	}

	err = u.serviceRepo.Update(c, req.CompanyID, req.ServiceID, su)
	if err != nil {
		if err == repository.ErrNoItems {
			return service.NotFoundError
		} else if err == repository.ErrInvalidParamInput {
			return service.NothingToUpdateError
		}
		return err
	}

	return nil
}

func (u *Usecase) Delete(ctx context.Context, req *dto.DeleteReq) error {
	if err := u.validator.RawRequest(req); err != nil {
		return service.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	_, err := u.companyRepo.Get(c, req.CompanyID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return company.NotFoundError
		}
		return err
	}

	err = u.tx.WithinTx(c, func(c context.Context) error {
		err := u.reservationRepo.DeleteByService(c, req.ServiceID)
		if err != nil {
			return err
		}

		return u.serviceRepo.Delete(c, req.CompanyID, req.ServiceID)
	})
	if err != nil {
		if err == repository.ErrNoItems {
			return service.NotFoundError
		}
		return err
	}

	return nil
}

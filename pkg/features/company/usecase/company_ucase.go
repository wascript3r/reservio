package usecase

import (
	"context"
	"html"
	"time"

	"github.com/wascript3r/reservio/pkg/features/reservation"

	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/company/dto"
	"github.com/wascript3r/reservio/pkg/features/company/models"
	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/features/user"
	"github.com/wascript3r/reservio/pkg/repository"
)

type Usecase struct {
	tx              repository.Transactor
	companyRepo     company.Repository
	serviceRepo     service.Repository
	reservationRepo reservation.Repository
	userRepo        user.Repository
	ctxTimeout      time.Duration

	validator company.Validator
	userUcase user.Usecase
}

func New(tx repository.Transactor, cr company.Repository, sr service.Repository, rs reservation.Repository, ur user.Repository, t time.Duration, v company.Validator, uu user.Usecase) *Usecase {
	return &Usecase{
		tx:              tx,
		companyRepo:     cr,
		serviceRepo:     sr,
		reservationRepo: rs,
		userRepo:        ur,
		ctxTimeout:      t,

		validator: v,
		userUcase: uu,
	}
}

func (u *Usecase) Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, company.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var companyID string
	err := u.tx.WithinTx(c, func(c context.Context) error {
		var err error
		companyID, err = u.userUcase.Create(c, &req.CreateReq, false)
		if err != nil {
			return err
		}

		req.Escape(false)
		cs := &models.Company{
			ID:          companyID,
			Name:        req.Name,
			Address:     req.Address,
			Description: req.Description,
		}

		return u.companyRepo.Insert(c, cs)
	})
	if err != nil {
		return nil, err
	}

	return &dto.CreateRes{
		ID: companyID,
	}, nil
}

func (u *Usecase) Get(ctx context.Context, req *dto.GetReq, onlyApproved bool) (*dto.GetRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, company.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ci, err := u.companyRepo.Get(c, req.CompanyID, onlyApproved)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, company.NotFoundError
		}
		return nil, err
	}

	return &dto.GetRes{
		ID:          ci.ID,
		Name:        ci.Name,
		Address:     ci.Address,
		Description: ci.Description,
		Email:       ci.Email,
		Approved:    ci.Approved,
	}, nil
}

func (u *Usecase) GetAll(ctx context.Context, onlyApproved bool) (*dto.GetAllRes, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cis, err := u.companyRepo.GetAll(c, onlyApproved)
	if err != nil {
		return nil, err
	}

	res := &dto.GetAllRes{
		Companies: make([]*dto.Company, len(cis)),
	}
	for i, ci := range cis {
		res.Companies[i] = &dto.Company{
			ID:          ci.ID,
			Name:        ci.Name,
			Address:     ci.Address,
			Description: ci.Description,
			Email:       ci.Email,
			Approved:    ci.Approved,
		}
	}

	return res, nil
}

func (u *Usecase) Update(ctx context.Context, req *dto.UpdateReq) error {
	if err := u.validator.RawRequest(req); err != nil {
		return company.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	if req.Name != nil {
		*req.Name = html.EscapeString(*req.Name)
	}
	if req.Address != nil {
		*req.Address = html.EscapeString(*req.Address)
	}
	if req.Description != nil {
		*req.Description = html.EscapeString(*req.Description)
	}

	err := u.companyRepo.Update(c, req.CompanyID, &models.CompanyUpdate{
		Name:        req.Name,
		Address:     req.Address,
		Description: req.Description,
		Approved:    req.Approved,
	})
	if err != nil {
		if err == repository.ErrNoItems {
			return company.NotFoundError
		} else if err == repository.ErrInvalidParamInput {
			return company.NothingToUpdateError
		}
		return err
	}

	return nil
}

func (u *Usecase) Delete(ctx context.Context, req *dto.DeleteReq) error {
	if err := u.validator.RawRequest(req); err != nil {
		return company.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	err := u.tx.WithinTx(c, func(c context.Context) error {
		err := u.reservationRepo.DeleteByCompany(c, req.CompanyID)
		if err != nil {
			return err
		}

		err = u.serviceRepo.DeleteByCompany(c, req.CompanyID)
		if err != nil {
			return err
		}

		err = u.companyRepo.Delete(c, req.CompanyID)
		if err != nil {
			return err
		}

		return u.userRepo.Delete(c, req.CompanyID)
	})
	if err != nil {
		if err == repository.ErrNoItems {
			return company.NotFoundError
		}
		return err
	}

	return nil
}

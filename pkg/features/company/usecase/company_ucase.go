package usecase

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/company/dto"
	"github.com/wascript3r/reservio/pkg/features/company/models"
	"github.com/wascript3r/reservio/pkg/features/user"
	"github.com/wascript3r/reservio/pkg/repository"
)

type Usecase struct {
	tx          repository.Transactor
	companyRepo company.Repository
	ctxTimeout  time.Duration

	validator company.Validator
	userUcase user.Usecase
}

func New(tx repository.Transactor, cr company.Repository, t time.Duration, v company.Validator, uu user.Usecase) *Usecase {
	return &Usecase{
		tx:          tx,
		companyRepo: cr,
		ctxTimeout:  t,

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

		cs := &models.Company{
			UserID:      companyID,
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

func (u *Usecase) GetAll(ctx context.Context) (*dto.GetAllRes, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cis, err := u.companyRepo.GetAll(c)
	if err != nil {
		return nil, err
	}

	res := &dto.GetAllRes{
		Companies: make([]*dto.Company, len(cis)),
	}
	for i, ci := range cis {
		res.Companies[i] = &dto.Company{
			ID:          ci.UserID,
			Name:        ci.Name,
			Address:     ci.Address,
			Description: ci.Description,
			Email:       ci.Email,
		}
	}

	return res, nil
}

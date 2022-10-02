package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	"github.com/wascript3r/reservio/pkg/features/reservation/dto"
	"github.com/wascript3r/reservio/pkg/features/reservation/models"
	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/repository"
)

const dateFormat = "2006-01-02 15:04"

type Usecase struct {
	reservationRepo reservation.Repository
	serviceRepo     service.Repository
	companyRepo     company.Repository
	ctxTimeout      time.Duration

	validator reservation.Validator
}

func New(rr reservation.Repository, sr service.Repository, cr company.Repository, t time.Duration, v reservation.Validator) *Usecase {
	return &Usecase{
		reservationRepo: rr,
		serviceRepo:     sr,
		companyRepo:     cr,
		ctxTimeout:      t,

		validator: v,
	}
}

func (u *Usecase) Create(ctx context.Context, req *dto.CreateReq) (*dto.CreateRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, reservation.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	date, err := time.Parse(dateFormat, req.Date)
	if err != nil {
		return nil, reservation.InvalidInputError
	}
	fmt.Println(date.Location(), date)

	_, err = u.companyRepo.Get(c, req.CompanyID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, company.NotFoundError
		}
		return nil, err
	}

	ss, err := u.serviceRepo.Get(c, req.CompanyID, req.ServiceID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, service.NotFoundError
		}
		return nil, err
	}

	if err := u.validator.ReservationDate(ss, date); err != nil {
		return nil, err
	}

	req.Escape()
	exists, err := u.reservationRepo.Exists(c, req.CompanyID, req.ServiceID, date)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, reservation.AlreadyExistsError
	}

	rs := &models.Reservation{
		ServiceID: req.ServiceID,
		Date:      date,
		Comment:   req.Comment,
	}

	id, err := u.reservationRepo.Insert(c, rs)
	if err != nil {
		if err == repository.ErrConflictWithRelatedItems {
			return nil, service.NotFoundError
		}
	}

	return &dto.CreateRes{
		ID: id,
	}, nil
}

func (u *Usecase) Get(ctx context.Context, req *dto.GetReq, onlyApprovedCompany bool) (*dto.GetRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, reservation.InvalidInputError.SetData(err.GetData())
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

	_, err = u.serviceRepo.Get(c, req.CompanyID, req.ServiceID, onlyApprovedCompany)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, service.NotFoundError
		}
		return nil, err
	}

	rs, err := u.reservationRepo.Get(c, req.CompanyID, req.ServiceID, req.ReservationID, onlyApprovedCompany)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, reservation.NotFoundError
		}
		return nil, err
	}

	return &dto.GetRes{
		ID:        rs.ID,
		ServiceID: rs.ServiceID,
		Date:      rs.Date.UTC().Format(dateFormat),
		Comment:   rs.Comment,
	}, nil
}

func (u *Usecase) GetAll(ctx context.Context, req *dto.GetAllReq, onlyApprovedCompany bool) (*dto.GetAllRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, reservation.InvalidInputError.SetData(err.GetData())
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

	_, err = u.serviceRepo.Get(c, req.CompanyID, req.ServiceID, onlyApprovedCompany)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, service.NotFoundError
		}
		return nil, err
	}

	rss, err := u.reservationRepo.GetAll(c, req.CompanyID, req.ServiceID, onlyApprovedCompany)
	if err != nil {
		return nil, err
	}

	res := make([]*dto.Reservation, len(rss))
	for i, rs := range rss {
		res[i] = &dto.Reservation{
			ID:        rs.ID,
			ServiceID: rs.ServiceID,
			Date:      rs.Date.UTC().Format(dateFormat),
			Comment:   rs.Comment,
		}
	}

	return &dto.GetAllRes{
		Reservations: res,
	}, nil
}

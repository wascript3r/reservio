package usecase

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/client"
	cldto "github.com/wascript3r/reservio/pkg/features/client/dto"
	"github.com/wascript3r/reservio/pkg/features/company"
	cdto "github.com/wascript3r/reservio/pkg/features/company/dto"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	"github.com/wascript3r/reservio/pkg/features/reservation/dto"
	"github.com/wascript3r/reservio/pkg/features/reservation/models"
	"github.com/wascript3r/reservio/pkg/features/service"
	sdto "github.com/wascript3r/reservio/pkg/features/service/dto"
	"github.com/wascript3r/reservio/pkg/repository"
)

const dateFormat = "2006-01-02 15:04"

type Usecase struct {
	reservationRepo reservation.Repository
	serviceRepo     service.Repository
	companyRepo     company.Repository
	clientRepo      client.Repository
	ctxTimeout      time.Duration

	validator reservation.Validator
}

func New(rr reservation.Repository, sr service.Repository, cr company.Repository, clr client.Repository, t time.Duration, v reservation.Validator) *Usecase {
	return &Usecase{
		reservationRepo: rr,
		serviceRepo:     sr,
		companyRepo:     cr,
		clientRepo:      clr,
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

	exists, err := u.reservationRepo.Exists(c, req.CompanyID, req.ServiceID, date)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, reservation.AlreadyExistsError
	}

	req.Escape()
	rs := &models.Reservation{
		ServiceID: req.ServiceID,
		ClientID:  req.ClientID,
		Date:      date,
		Comment:   req.Comment,
	}

	id, err := u.reservationRepo.Insert(c, rs)
	if err != nil {
		if err == repository.ErrConflictWithRelatedItems {
			return nil, client.NotFoundError
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
		Client: &cldto.Client{
			ID:        rs.Client.ID,
			FirstName: rs.Client.FirstName,
			LastName:  rs.Client.LastName,
			Phone:     rs.Client.Phone,
			Email:     rs.Client.Email,
		},
		Date:     rs.Date.UTC().Format(dateFormat),
		Comment:  rs.Comment,
		Approved: rs.Approved,
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
			Client: &cldto.Client{
				ID:        rs.Client.ID,
				FirstName: rs.Client.FirstName,
				LastName:  rs.Client.LastName,
				Phone:     rs.Client.Phone,
				Email:     rs.Client.Email,
			},
			Date:     rs.Date.UTC().Format(dateFormat),
			Comment:  rs.Comment,
			Approved: rs.Approved,
		}
	}

	return &dto.GetAllRes{
		Reservations: res,
	}, nil
}

func (u *Usecase) GetAllByClient(ctx context.Context, req *dto.GetAllByClientReq) (*dto.GetAllByClientRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, reservation.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	_, err := u.clientRepo.Get(c, req.ClientID)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, client.NotFoundError
		}
		return nil, err
	}

	rss, err := u.reservationRepo.GetAllByClient(c, req.ClientID)
	if err != nil {
		return nil, err
	}

	res := make([]*dto.ClientReservation, len(rss))
	for i, rs := range rss {
		ws := make(sdto.WorkSchedule)
		for k, v := range rs.Service.WorkSchedule {
			ws[k] = (*sdto.WorkHours)(v)
		}

		res[i] = &dto.ClientReservation{
			ID: rs.ID,
			Service: &sdto.FullService{
				ID: rs.Service.ID,
				Company: &cdto.Company{
					ID:          rs.Service.Company.ID,
					Email:       rs.Service.Company.Email,
					Name:        rs.Service.Company.Name,
					Address:     rs.Service.Company.Address,
					Description: rs.Service.Company.Description,
					Approved:    rs.Service.Company.Approved,
				},
				Title:           rs.Service.Title,
				Description:     rs.Service.Description,
				SpecialistName:  rs.Service.SpecialistName,
				SpecialistPhone: rs.Service.SpecialistPhone,
				VisitDuration:   rs.Service.VisitDuration,
				WorkSchedule:    ws,
			},
			Date:     rs.Date.UTC().Format(dateFormat),
			Comment:  rs.Comment,
			Approved: rs.Approved,
		}
	}

	return &dto.GetAllByClientRes{
		Reservations: res,
	}, nil
}

func (u *Usecase) Update(ctx context.Context, req *dto.UpdateReq) error {
	if err := u.validator.RawRequest(req); err != nil {
		return reservation.InvalidInputError.SetData(err.GetData())
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var date *time.Time
	if req.Date != nil {
		d, err := time.Parse(dateFormat, *req.Date)
		if err != nil {
			return reservation.InvalidInputError
		}
		date = &d
	}

	_, err := u.companyRepo.Get(c, req.CompanyID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return company.NotFoundError
		}
		return err
	}

	ss, err := u.serviceRepo.Get(c, req.CompanyID, req.ServiceID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return service.NotFoundError
		}
		return err
	}

	rs, err := u.reservationRepo.Get(c, req.CompanyID, req.ServiceID, req.ReservationID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return reservation.NotFoundError
		}
		return err
	}

	if date != nil && !date.Equal(rs.Date) {
		if err := u.validator.ReservationDate(ss, *date); err != nil {
			return err
		}

		exists, err := u.reservationRepo.Exists(c, req.CompanyID, req.ServiceID, *date)
		if err != nil {
			return err
		} else if exists {
			return reservation.AlreadyExistsError
		}
	}

	req.Escape()
	ru := &models.ReservationUpdate{
		Date:     date,
		Approved: req.Approved,
	}
	if req.Comment != nil {
		ru.Comment = &req.Comment.Value
	}

	err = u.reservationRepo.Update(c, req.CompanyID, req.ServiceID, req.ReservationID, ru)
	if err != nil {
		if err == repository.ErrNoItems {
			return reservation.NotFoundError
		}
	}

	return nil
}

func (u *Usecase) Delete(ctx context.Context, req *dto.DeleteReq) error {
	if err := u.validator.RawRequest(req); err != nil {
		return reservation.InvalidInputError.SetData(err.GetData())
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

	_, err = u.serviceRepo.Get(c, req.CompanyID, req.ServiceID, false)
	if err != nil {
		if err == repository.ErrNoItems {
			return service.NotFoundError
		}
		return err
	}

	err = u.reservationRepo.Delete(c, req.CompanyID, req.ServiceID, req.ReservationID)
	if err != nil {
		if err == repository.ErrNoItems {
			return reservation.NotFoundError
		}
	}

	return nil
}

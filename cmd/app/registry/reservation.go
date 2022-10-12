package registry

import (
	"context"

	"github.com/julienschmidt/httprouter"
	httpHnd "github.com/wascript3r/reservio/pkg/features/reservation/delivery/http"
	"github.com/wascript3r/reservio/pkg/features/reservation/repository"
	"github.com/wascript3r/reservio/pkg/features/reservation/usecase"
	"github.com/wascript3r/reservio/pkg/features/reservation/validator"
)

type ReservationReg struct {
	*GlobalReg

	repository *repository.PgRepo
	validator  *validator.Validator
	usecase    *usecase.Usecase
}

func NewReservation(gr *GlobalReg) *ReservationReg {
	return &ReservationReg{
		GlobalReg: gr,
	}
}

func (r *ReservationReg) Repository() *repository.PgRepo {
	if r.repository == nil {
		r.repository = repository.NewPgRepo(r.db)
	}

	return r.repository
}

func (r *ReservationReg) Validator() *validator.Validator {
	if r.validator == nil {
		r.validator = validator.New()
	}

	return r.validator
}

func (r *ReservationReg) Usecase() *usecase.Usecase {
	if r.usecase == nil {
		r.usecase = usecase.New(
			r.Repository(),
			r.serviceReg.Repository(),
			r.companyReg.Repository(),
			r.clientReg.Repository(),
			r.cfg.Database.Postgres.QueryTimeout.Duration,

			r.Validator(),
		)
	}

	return r.usecase
}

func (r *ReservationReg) ServeHTTP(router *httprouter.Router) {
	httpHnd.NewHTTPHandler(
		context.Background(),
		router,
		r.tokenReg.ClientMid(),
		r.tokenReg.CompanyOrClientMid(),
		r.tokenReg.ParseMid(),

		r.mapper,
		r.Usecase(),
		r.tokenReg.Usecase(),
	)
}

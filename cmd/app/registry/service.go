package registry

import (
	"github.com/julienschmidt/httprouter"
	httpHnd "github.com/wascript3r/reservio/pkg/features/service/delivery/http"
	"github.com/wascript3r/reservio/pkg/features/service/repository"
	"github.com/wascript3r/reservio/pkg/features/service/usecase"
	"github.com/wascript3r/reservio/pkg/features/service/validator"
)

type ServiceReg struct {
	*GlobalReg

	repository *repository.PgRepo
	validator  *validator.Validator
	usecase    *usecase.Usecase
}

func NewService(gr *GlobalReg) *ServiceReg {
	return &ServiceReg{
		GlobalReg: gr,
	}
}

func (r *ServiceReg) Repository() *repository.PgRepo {
	if r.repository == nil {
		r.repository = repository.NewPgRepo(r.db)
	}

	return r.repository
}

func (r *ServiceReg) Validator() *validator.Validator {
	if r.validator == nil {
		r.validator = validator.New()
	}

	return r.validator
}

func (r *ServiceReg) Usecase() *usecase.Usecase {
	if r.usecase == nil {
		r.usecase = usecase.New(
			r.db,
			r.Repository(),
			r.reservationReg.Repository(),
			r.companyReg.Repository(),
			r.cfg.Database.Postgres.QueryTimeout.Duration,

			r.Validator(),
		)
	}

	return r.usecase
}

func (r *ServiceReg) ServeHTTP(router *httprouter.Router) {
	httpHnd.NewHTTPHandler(router, r.mapper, r.Usecase())
}

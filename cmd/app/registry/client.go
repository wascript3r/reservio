package registry

import (
	"github.com/julienschmidt/httprouter"
	httpHnd "github.com/wascript3r/reservio/pkg/features/client/delivery/http"
	"github.com/wascript3r/reservio/pkg/features/client/repository"
	"github.com/wascript3r/reservio/pkg/features/client/usecase"
	"github.com/wascript3r/reservio/pkg/features/client/validator"
)

type ClientReg struct {
	*GlobalReg

	repository *repository.PgRepo
	validator  *validator.Validator
	usecase    *usecase.Usecase
}

func NewClient(gr *GlobalReg) *ClientReg {
	return &ClientReg{
		GlobalReg: gr,
	}
}

func (r *ClientReg) Repository() *repository.PgRepo {
	if r.repository == nil {
		r.repository = repository.NewPgRepo(r.db)
	}

	return r.repository
}

func (r *ClientReg) Validator() *validator.Validator {
	if r.validator == nil {
		r.validator = validator.New(r.userReg.Validator())
	}

	return r.validator
}

func (r *ClientReg) Usecase() *usecase.Usecase {
	if r.usecase == nil {
		r.usecase = usecase.New(
			r.db,
			r.Repository(),
			r.cfg.Database.Postgres.QueryTimeout.Duration,

			r.Validator(),
			r.userReg.Usecase(),
		)
	}

	return r.usecase
}

func (r *ClientReg) ServeHTTP(router *httprouter.Router) {
	httpHnd.NewHTTPHandler(router, r.mapper, r.Usecase(), r.reservationReg.Usecase())
}

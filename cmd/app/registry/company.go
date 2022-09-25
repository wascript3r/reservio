package registry

import (
	"github.com/julienschmidt/httprouter"
	httpHnd "github.com/wascript3r/reservio/pkg/features/company/delivery/http"
	"github.com/wascript3r/reservio/pkg/features/company/repository"
	"github.com/wascript3r/reservio/pkg/features/company/usecase"
	"github.com/wascript3r/reservio/pkg/features/company/validator"
)

type CompanyReg struct {
	*GlobalReg

	repository *repository.PgRepo
	validator  *validator.Validator
	usecase    *usecase.Usecase
}

func NewCompany(gr *GlobalReg) *CompanyReg {
	return &CompanyReg{
		GlobalReg: gr,
	}
}

func (r *CompanyReg) Repository() *repository.PgRepo {
	if r.repository == nil {
		r.repository = repository.NewPgRepo(r.db)
	}

	return r.repository
}

func (r *CompanyReg) Validator() *validator.Validator {
	if r.validator == nil {
		r.validator = validator.New(r.userReg.Validator())
	}

	return r.validator
}

func (r *CompanyReg) Usecase() *usecase.Usecase {
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

func (r *CompanyReg) ServeHTTP(router *httprouter.Router) {
	httpHnd.NewHTTPHandler(router, r.Usecase())
}

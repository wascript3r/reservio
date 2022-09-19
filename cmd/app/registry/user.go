package registry

import (
	"github.com/julienschmidt/httprouter"
	httpHnd "github.com/wascript3r/reservio/pkg/features/user/delivery/http"
	"github.com/wascript3r/reservio/pkg/features/user/pwhasher"
	"github.com/wascript3r/reservio/pkg/features/user/repository"
	"github.com/wascript3r/reservio/pkg/features/user/usecase"
	"github.com/wascript3r/reservio/pkg/features/user/validator"
)

type UserReg struct {
	*GlobalReg

	repository *repository.PgRepo
	pwHasher   *pwhasher.PwHasher
	validator  *validator.Validator
	usecase    *usecase.Usecase
}

func NewUser(gr *GlobalReg) *UserReg {
	return &UserReg{
		GlobalReg: gr,
	}
}

func (r *UserReg) Repository() *repository.PgRepo {
	if r.repository == nil {
		r.repository = repository.NewPgRepo(r.db)
	}

	return r.repository
}

func (r *UserReg) PwHasher() *pwhasher.PwHasher {
	if r.pwHasher == nil {
		pwh := pwhasher.New(r.cfg.Auth.PasswordCost)
		r.pwHasher = &pwh
	}

	return r.pwHasher
}

func (r *UserReg) Validator() *validator.Validator {
	if r.validator == nil {
		r.validator = validator.New(r.Repository())
	}

	return r.validator
}

func (r *UserReg) Usecase() *usecase.Usecase {
	if r.usecase == nil {
		r.usecase = usecase.New(
			r.Repository(),
			r.cfg.Database.Postgres.QueryTimeout.Duration,

			r.PwHasher(),
			r.Validator(),
		)
	}

	return r.usecase
}

func (r *UserReg) ServeHTTP(router *httprouter.Router) {
	httpHnd.NewHTTPHandler(router, r.Usecase())
}

package registry

import (
	"github.com/julienschmidt/httprouter"
	stack "github.com/wascript3r/httputil/middleware"
	"github.com/wascript3r/reservio/pkg/features/token/delivery/http"
	httpHnd "github.com/wascript3r/reservio/pkg/features/token/delivery/http"
	mid "github.com/wascript3r/reservio/pkg/features/token/delivery/http/middleware"
	"github.com/wascript3r/reservio/pkg/features/token/jwt"
	"github.com/wascript3r/reservio/pkg/features/token/repository"
	"github.com/wascript3r/reservio/pkg/features/token/validator"
	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

type TokenReg struct {
	*GlobalReg

	repository *repository.PgRepo
	validator  *validator.Validator
	usecase    *jwt.Usecase
	mid        *mid.HTTPMiddleware

	authMid            *http.Auth
	parseMid           *http.Parse
	adminMid           *http.Admin
	companyMid         *http.Company
	clientMid          *http.Client
	companyOrClientMid *http.CompanyOrClient
}

func NewToken(gr *GlobalReg) *TokenReg {
	return &TokenReg{
		GlobalReg: gr,
	}
}

func (r *TokenReg) Repository() *repository.PgRepo {
	if r.repository == nil {
		r.repository = repository.NewPgRepo(r.db)
	}

	return r.repository
}

func (r *TokenReg) Validator() *validator.Validator {
	if r.validator == nil {
		r.validator = validator.New()
	}

	return r.validator
}

func (r *TokenReg) Usecase() *jwt.Usecase {
	if r.usecase == nil {
		r.usecase = jwt.New(
			&jwt.Options{
				PrivateKey:        []byte(r.cfg.Auth.JWT.SecretKey),
				AccessExpiration:  r.cfg.Auth.JWT.AccessExpiration.Duration,
				RefreshExpiration: r.cfg.Auth.JWT.RefreshExpiration.Duration,
				Issuer:            r.cfg.Auth.JWT.Issuer,
			},

			r.db,
			r.Repository(),
			r.cfg.Database.Postgres.QueryTimeout.Duration,
			r.Validator(),
		)
	}

	return r.usecase
}

func (r *TokenReg) HTTPMid() *mid.HTTPMiddleware {
	if r.mid == nil {
		r.mid = mid.NewHTTPMiddleware(r.Usecase())
	}

	return r.mid
}

func (r *TokenReg) AuthMid() http.Auth {
	if r.authMid == nil {
		r.authMid = &http.Auth{StackCtx: stack.NewCtx()}
		r.authMid.Use(r.HTTPMid().Authenticated)
	}

	return *r.authMid
}

func (r *TokenReg) ParseMid() http.Parse {
	if r.parseMid == nil {
		r.parseMid = &http.Parse{StackCtx: stack.NewCtx()}
		r.parseMid.Use(r.HTTPMid().ParseUser)
	}

	return *r.parseMid
}

func (r *TokenReg) AdminMid() http.Admin {
	if r.adminMid == nil {
		r.adminMid = &http.Admin{StackCtx: stack.NewCtx()}
		r.adminMid.Use(r.HTTPMid().IsOneOf(umodels.AdminRole))
	}

	return *r.adminMid
}

func (r *TokenReg) CompanyMid() http.Company {
	if r.companyMid == nil {
		r.companyMid = &http.Company{StackCtx: stack.NewCtx()}
		r.companyMid.Use(r.HTTPMid().IsOneOf(umodels.CompanyRole))
	}

	return *r.companyMid
}

func (r *TokenReg) ClientMid() http.Client {
	if r.clientMid == nil {
		r.clientMid = &http.Client{StackCtx: stack.NewCtx()}
		r.clientMid.Use(r.HTTPMid().IsOneOf(umodels.ClientRole))
	}

	return *r.clientMid
}

func (r *TokenReg) CompanyOrClientMid() http.CompanyOrClient {
	if r.companyOrClientMid == nil {
		r.companyOrClientMid = &http.CompanyOrClient{StackCtx: stack.NewCtx()}
		r.companyOrClientMid.Use(r.HTTPMid().IsOneOf(umodels.CompanyRole, umodels.ClientRole))
	}

	return *r.companyOrClientMid
}

func (r *TokenReg) ServeHTTP(router *httprouter.Router) {
	httpHnd.NewHTTPHandler(
		router,

		r.mapper,
		r.Usecase(),
	)
}

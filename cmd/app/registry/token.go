package registry

import (
	stack "github.com/wascript3r/httputil/middleware"
	"github.com/wascript3r/reservio/pkg/features/token/delivery/http"
	mid "github.com/wascript3r/reservio/pkg/features/token/delivery/http/middleware"
	"github.com/wascript3r/reservio/pkg/features/token/jwt"
	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

type TokenReg struct {
	*GlobalReg

	usecase *jwt.Usecase
	mid     *mid.HTTPMiddleware

	authMid    *http.Auth
	adminMid   *http.Admin
	companyMid *http.Company
	clientMid  *http.Client
}

func NewToken(gr *GlobalReg) *TokenReg {
	return &TokenReg{
		GlobalReg: gr,
	}
}

func (r *TokenReg) Usecase() *jwt.Usecase {
	if r.usecase == nil {
		r.usecase = jwt.New(
			[]byte(r.cfg.Auth.JWT.SecretKey),
			r.cfg.Auth.JWT.Expiration.Duration,
			r.cfg.Auth.JWT.Issuer,
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

func (r *TokenReg) AdminMid() http.Admin {
	if r.adminMid == nil {
		r.adminMid = &http.Admin{StackCtx: stack.NewCtx()}
		r.adminMid.Use(r.HTTPMid().HasRole(umodels.AdminRole))
	}

	return *r.adminMid
}

func (r *TokenReg) CompanyMid() http.Company {
	if r.companyMid == nil {
		r.companyMid = &http.Company{StackCtx: stack.NewCtx()}
		r.companyMid.Use(r.HTTPMid().HasRole(umodels.CompanyRole))
	}

	return *r.companyMid
}

func (r *TokenReg) ClientMid() http.Client {
	if r.clientMid == nil {
		r.clientMid = &http.Client{StackCtx: stack.NewCtx()}
		r.clientMid.Use(r.HTTPMid().HasRole(umodels.ClientRole))
	}

	return *r.clientMid
}

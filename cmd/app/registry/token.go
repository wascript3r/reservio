package registry

import (
	"github.com/wascript3r/reservio/pkg/features/token/jwt"
)

type TokenReg struct {
	*GlobalReg

	usecase *jwt.Usecase
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

package registry

import (
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/cmd/app/config"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

type GlobalReg struct {
	cfg    *config.Config
	db     *pgsql.Database
	mapper *httpjson.CodeMapper

	userReg        *UserReg
	companyReg     *CompanyReg
	serviceReg     *ServiceReg
	reservationReg *ReservationReg
	clientReg      *ClientReg
	tokenReg       *TokenReg
	loggerReg      *LoggerReg
}

func NewGlobal(cfg *config.Config, db *pgsql.Database, cm *httpjson.CodeMapper) *GlobalReg {
	return &GlobalReg{
		cfg:    cfg,
		db:     db,
		mapper: cm,
	}
}

func (r *GlobalReg) SetUserReg(userReg *UserReg) *GlobalReg {
	r.userReg = userReg
	return r
}

func (r *GlobalReg) SetCompanyReg(companyReg *CompanyReg) *GlobalReg {
	r.companyReg = companyReg
	return r
}

func (r *GlobalReg) SetServiceReg(serviceReg *ServiceReg) *GlobalReg {
	r.serviceReg = serviceReg
	return r
}

func (r *GlobalReg) SetReservationReg(reservationReg *ReservationReg) *GlobalReg {
	r.reservationReg = reservationReg
	return r
}

func (r *GlobalReg) SetClientReg(clientReg *ClientReg) *GlobalReg {
	r.clientReg = clientReg
	return r
}

func (r *GlobalReg) SetTokenReg(tokenReg *TokenReg) *GlobalReg {
	r.tokenReg = tokenReg
	return r
}

func (r *GlobalReg) SetLoggerReg(loggerReg *LoggerReg) *GlobalReg {
	r.loggerReg = loggerReg
	return r
}

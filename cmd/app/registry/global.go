package registry

import (
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

type GlobalReg struct {
	cfg *Config
	db  *pgsql.Database

	userReg    *UserReg
	companyReg *CompanyReg
	loggerReg  *LoggerReg
}

func NewGlobal(cfg *Config, db *pgsql.Database) *GlobalReg {
	return &GlobalReg{
		cfg: cfg,
		db:  db,
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

func (r *GlobalReg) SetLoggerReg(loggerReg *LoggerReg) *GlobalReg {
	r.loggerReg = loggerReg
	return r
}

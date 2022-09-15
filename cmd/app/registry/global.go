package registry

import (
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

type GlobalReg struct {
	cfg *Config
	db  *pgsql.Database

	loggerReg *LoggerReg
}

func NewGlobal(cfg *Config, db *pgsql.Database) *GlobalReg {
	return &GlobalReg{
		cfg: cfg,
		db:  db,
	}
}

func (r *GlobalReg) SetLoggerReg(loggerReg *LoggerReg) *GlobalReg {
	r.loggerReg = loggerReg
	return r
}

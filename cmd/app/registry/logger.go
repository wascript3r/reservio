package registry

import (
	"log"
	"os"

	"github.com/wascript3r/reservio/pkg/logger/usecase"
)

const AppLoggerPrefix = "[APP]"

type LoggerReg struct {
	cfg     *Config
	usecase *usecase.Usecase
}

func NewLogger(cfg *Config) *LoggerReg {
	return &LoggerReg{
		cfg: cfg,
	}
}

func (r *LoggerReg) Usecase() *usecase.Usecase {
	if r.usecase == nil {
		logFlags := 0
		if r.cfg.Log.ShowTimestamp {
			logFlags = log.Ltime
		}
		r.usecase = usecase.New(
			AppLoggerPrefix,
			log.New(os.Stdout, "", logFlags),
		)
	}

	return r.usecase
}

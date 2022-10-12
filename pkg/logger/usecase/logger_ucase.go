package usecase

import (
	"log"
)

type Usecase struct {
	prefix string
	log    *log.Logger
}

func New(prefix string, log *log.Logger) *Usecase {
	return &Usecase{prefix, log}
}

func (l *Usecase) Info(format string, v ...any) {
	v = append([]any{l.prefix}, v...)
	l.log.Printf("[INFO] %s: "+format, v...)
}

func (l *Usecase) Error(format string, v ...any) {
	v = append([]any{l.prefix}, v...)
	l.log.Printf("[ERROR] %s: "+format, v...)
}

package pgsql

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/wascript3r/reservio/pkg/repository"
)

type ErrCode string

const (
	UniqueViolationErrCode ErrCode = "23505"
	CheckViolationErrCode  ErrCode = "23514"
)

func ParseReadErr(err error) error {
	if err == nil {
		return nil
	}

	switch err {
	case sql.ErrNoRows:
		return repository.ErrNoItems
	}
	return err
}

func ParseWriteErr(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*pq.Error); ok {
		switch ErrCode(e.Code) {
		case UniqueViolationErrCode:
			return repository.ErrItemExists

		case CheckViolationErrCode:
			return repository.ErrIntegrityViolation
		}
	}
	return err
}

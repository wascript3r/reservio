package errutil

import (
	"net/http"

	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
)

type ErrType uint8

const (
	InvalidInputErrType ErrType = iota + 1
	UnkownErrType
	OtherErrType
)

func ParseErrFunc(invalidInputErr, unknownErr *errcode.Error) func(error) (ErrType, *errcode.Error) {
	return func(err error) (ErrType, *errcode.Error) {
		if err == invalidInputErr {
			return InvalidInputErrType, invalidInputErr
		}

		code := errcode.UnwrapErr(err, unknownErr)
		if code == unknownErr {
			return UnkownErrType, unknownErr
		}

		return OtherErrType, code
	}
}

func ServeHTTP(w http.ResponseWriter, et ErrType, code *errcode.Error) {
	switch et {
	case InvalidInputErrType:
		httpjson.BadRequestCustom(w, code, nil)
	case UnkownErrType:
		httpjson.InternalErrorCustom(w, code, nil)
	default:
		httpjson.ServeErr(w, code, nil)
	}
}

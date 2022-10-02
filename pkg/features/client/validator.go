package client

import (
	"github.com/wascript3r/reservio/pkg/validator"
)

type Validator interface {
	RawRequest(s any) validator.Error
}

package user

import (
	"errors"

	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	ErrEmailExists = errors.New("email already exists")

	// Error codes

	InvalidInputError = errcode.InvalidInputError
	UnknownError      = errcode.UnknownError

	EmailAlreadyExistsError = errcode.New(
		"email_already_exists",
		errors.New("email already exists"),
	)
)

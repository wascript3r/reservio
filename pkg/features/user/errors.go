package user

import (
	"errors"

	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	// Error codes

	InvalidInputError = errcode.InvalidInputError
	UnknownError      = errcode.UnknownError

	EmailAlreadyExistsError = errcode.New(
		"email_already_exists",
		errors.New("email already exists"),
	)

	InvalidCredentialsError = errcode.New(
		"invalid_credentials",
		errors.New("invalid credentials"),
	)
)

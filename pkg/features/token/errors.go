package token

import (
	"errors"

	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	ErrCannotLoadToken = errors.New("cannot load token")

	// Error codes

	InvalidInputError = errcode.InvalidInputError
	UnknownError      = errcode.UnknownError

	InvalidOrExpiredTokenError = errcode.New(
		"token_invalid_or_expired",
		errors.New("token is invalid or expired"),
	)

	FaultyTokenError = errcode.New(
		"faulty_token",
		errors.New("faulty token provided"),
	)
)

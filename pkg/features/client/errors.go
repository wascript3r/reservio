package client

import (
	"errors"

	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	// Error codes

	InvalidInputError = errcode.InvalidInputError
	UnknownError      = errcode.UnknownError

	NotFoundError = errcode.New(
		"client_not_found",
		errors.New("client not found"),
	)
)

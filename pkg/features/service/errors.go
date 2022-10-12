package service

import (
	"errors"

	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	// Error codes

	InvalidInputError = errcode.InvalidInputError
	UnknownError      = errcode.UnknownError

	NotFoundError = errcode.New(
		"service_not_found",
		errors.New("service not found"),
	)

	NothingToUpdateError = errcode.New(
		"nothing_to_update",
		errors.New("nothing to update"),
	)
)

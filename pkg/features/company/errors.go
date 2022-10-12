package company

import (
	"errors"

	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	// Error codes

	InvalidInputError = errcode.InvalidInputError
	UnknownError      = errcode.UnknownError

	NotFoundError = errcode.New(
		"company_not_found",
		errors.New("company not found"),
	)

	NothingToUpdateError = errcode.New(
		"nothing_to_update",
		errors.New("nothing to update"),
	)
)

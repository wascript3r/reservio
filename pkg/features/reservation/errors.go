package reservation

import (
	"errors"

	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	// Error codes

	InvalidInputError = errcode.InvalidInputError
	UnknownError      = errcode.UnknownError

	NotFoundError = errcode.New(
		"reservation_not_found",
		errors.New("reservation not found"),
	)

	NothingToUpdateError = errcode.New(
		"nothing_to_update",
		errors.New("nothing to update"),
	)

	DateIsInPastError = errcode.New(
		"time_is_in_past",
		errors.New("specified reservation time is in the past"),
	)

	InvalidTimeError = errcode.New(
		"invalid_time",
		errors.New("invalid reservation time"),
	)

	ServiceNotAvailableError = errcode.New(
		"service_not_available",
		errors.New("service is not available at the specified time"),
	)

	AlreadyExistsError = errcode.New(
		"reservation_already_exists",
		errors.New("reservation already exists at the specified time"),
	)
)

package repository

import "errors"

var (
	ErrNoItems            = errors.New("no items found")
	ErrNullValue          = errors.New("item value is null")
	ErrItemExists         = errors.New("item already exists")
	ErrIntegrityViolation = errors.New("item violates integrity constraint")
	ErrInvalidParamInput  = errors.New("invalid input parameter")
)

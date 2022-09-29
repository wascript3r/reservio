package gov

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	tagError = map[string]func(fe validator.FieldError) string{
		"required": func(_ validator.FieldError) string {
			return "is required"
		},
		"gte": func(fe validator.FieldError) string {
			return fmt.Sprintf("must be greater than or equal to %s characters", fe.Param())
		},
		"gt": func(fe validator.FieldError) string {
			if fe.Param() == "0" {
				return "must be set"
			}
			return fmt.Sprintf("must be greater than %s characters", fe.Param())
		},
		"lte": func(fe validator.FieldError) string {
			return fmt.Sprintf("must be less than or equal to %s characters", fe.Param())
		},
		"lt": func(fe validator.FieldError) string {
			return fmt.Sprintf("must be less than %s characters", fe.Param())
		},
		"email": func(_ validator.FieldError) string {
			return "must be a valid email"
		},
		"uuid": func(_ validator.FieldError) string {
			return "must be a valid id"
		},
		"s_work_schedule": func(_ validator.FieldError) string {
			return "must contain valid week days"
		},
		"s_time": func(_ validator.FieldError) string {
			return "must be a valid time"
		},
		"e164": func(_ validator.FieldError) string {
			return "must be a valid phone number"
		},
	}
)

type Error struct {
	original error
	data     map[string]string
}

func (e *Error) Error() string {
	return e.original.Error()
}

func (e *Error) GetData() any {
	return e.data
}

func Translate(err error) *Error {
	ve, ok := err.(validator.ValidationErrors)
	if !ok || len(ve) == 0 {
		return &Error{err, nil}
	}

	data := make(map[string]string, len(ve))
	for _, fe := range ve {
		tagFn, ok := tagError[fe.ActualTag()]
		if !ok {
			continue
		}
		field := strings.Join(strings.Split(fe.StructNamespace(), ".")[1:], ".")
		data[field] = fmt.Sprintf("%s %s", fe.Field(), tagFn(fe))
	}

	return &Error{
		original: err,
		data:     data,
	}
}

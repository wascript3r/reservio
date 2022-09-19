package validator

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/features/user"
)

type Validator struct {
	govalidate    *validator.Validate
	userValidator user.SharedValidator
}

func New(uv user.SharedValidator) *Validator {
	v := validator.New()

	r := newRules()
	r.attachTo(v)
	r.attachExtTo(uv.GetRules(), v)

	return &Validator{
		govalidate:    v,
		userValidator: uv,
	}
}

func (v *Validator) RawRequest(s any) error {
	return v.govalidate.Struct(s)
}

func (v *Validator) EmailUniqueness(ctx context.Context, email string) error {
	return v.userValidator.EmailUniqueness(ctx, email)
}

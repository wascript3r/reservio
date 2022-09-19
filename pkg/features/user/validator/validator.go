package validator

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/features/user"
)

type Validator struct {
	govalidate *validator.Validate
	userRepo   user.Repository
}

func New(ur user.Repository) *Validator {
	v := validator.New()

	r := newRules()
	r.attachTo(v)

	return &Validator{v, ur}
}

func (v *Validator) RawRequest(s any) error {
	return v.govalidate.Struct(s)
}

func (v *Validator) EmailUniqueness(ctx context.Context, email string) error {
	exists, err := v.userRepo.EmailExists(ctx, email)
	if err != nil {
		return err
	}

	if exists {
		return user.ErrEmailExists
	}

	return nil
}

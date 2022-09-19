package validator

import (
	"context"

	gvalidator "github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/features/user"
	"github.com/wascript3r/reservio/pkg/validator"
)

type Validator struct {
	govalidate *gvalidator.Validate
	rules      rules
	userRepo   user.Repository
}

func New(ur user.Repository) *Validator {
	v := gvalidator.New()

	r := newRules()
	r.attachTo(v)

	return &Validator{
		govalidate: v,
		rules:      r,
		userRepo:   ur,
	}
}

func (v *Validator) RawRequest(s any) error {
	return v.govalidate.Struct(s)
}

func (v *Validator) GetRules() []validator.Rule {
	rs := make([]validator.Rule, len(v.rules.aliases))
	for i, a := range v.rules.aliases {
		rs[i] = a
	}
	return rs
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

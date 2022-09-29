package validator

import (
	gvalidator "github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/validator"
	"github.com/wascript3r/reservio/pkg/validator/gov"
)

type Validator struct {
	govalidate *gvalidator.Validate
}

func New() *Validator {
	v := gvalidator.New()

	r := newRules()
	r.attachTo(v)

	return &Validator{v}
}

func (v *Validator) RawRequest(s any) validator.Error {
	err := v.govalidate.Struct(s)
	if err != nil {
		return gov.Translate(err)
	}
	return nil
}

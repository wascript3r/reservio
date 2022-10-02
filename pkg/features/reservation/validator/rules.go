package validator

import (
	"reflect"
	"time"

	gvalidator "github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/validator/gov"
)

const dateFormat = "2006-01-02 15:04"

type rules struct {
	aliases []gov.AliasRule
	fns     []gov.FnRule
}

func newRules() rules {
	return rules{
		aliases: []gov.AliasRule{
			gov.NewAliasRule("r_comment", "gte=5,lte=200"),
		},
		fns: []gov.FnRule{
			gov.NewFnRule("r_date", validateDate),
		},
	}
}

func (r rules) attachTo(v *gvalidator.Validate) {
	for _, a := range r.aliases {
		v.RegisterAlias(a.Alias(), a.Tags())
	}
	for _, f := range r.fns {
		v.RegisterValidation(f.Name(), f.IsValid)
	}
}

func validateDate(fl gvalidator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() != reflect.String {
		return false
	}
	t := field.String()

	_, err := time.Parse(dateFormat, t)
	return err == nil
}

package validator

import (
	gvalidator "github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/validator"
	"github.com/wascript3r/reservio/pkg/validator/gov"
)

type rules struct {
	aliases []gov.AliasRule
}

func newRules() rules {
	return rules{
		aliases: []gov.AliasRule{
			gov.NewAliasRule("c_first_name", "gte=3,lte=50"),
			gov.NewAliasRule("c_last_name", "gte=3,lte=50"),
			gov.NewAliasRule("c_phone", "e164"),
		},
	}
}

func (r rules) attachTo(v *gvalidator.Validate) {
	for _, a := range r.aliases {
		v.RegisterAlias(a.Alias(), a.Tags())
	}
}

func (r rules) attachExtTo(rs []validator.Rule, v *gvalidator.Validate) {
	for _, rr := range rs {
		v.RegisterAlias(rr.Alias(), rr.Tags())
	}
}

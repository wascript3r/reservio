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
			gov.NewAliasRule("c_name", "gte=3,lte=100"),
			gov.NewAliasRule("c_address", "gte=5,lte=200"),
			gov.NewAliasRule("c_description", "gte=5"),
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

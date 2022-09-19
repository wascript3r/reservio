package validator

import (
	gvalidator "github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/validator/gov"
)

type rules struct {
	aliases []gov.AliasRule
}

func newRules() rules {
	return rules{
		aliases: []gov.AliasRule{
			gov.NewAliasRule("u_email", "lte=200,email"),
			gov.NewAliasRule("u_password", "gte=8,lte=100"),
		},
	}
}

func (r rules) attachTo(v *gvalidator.Validate) {
	for _, a := range r.aliases {
		v.RegisterAlias(a.Alias(), a.Tags())
	}
}

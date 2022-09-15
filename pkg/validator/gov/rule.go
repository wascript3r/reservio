package gov

import "github.com/go-playground/validator/v10"

type FnRule struct {
	name string
	fn   validator.Func
}

func NewFnRule(name string, fn validator.Func) FnRule {
	return FnRule{
		name: name,
		fn:   fn,
	}
}

func (r FnRule) Name() string {
	return r.name
}

func (r FnRule) IsValid(fl validator.FieldLevel) bool {
	return r.fn(fl)
}

type AliasRule struct {
	alias string
	tags  string
}

func NewAliasRule(alias, tags string) AliasRule {
	return AliasRule{
		alias: alias,
		tags:  tags,
	}
}

func (r AliasRule) Alias() string {
	return r.alias
}

func (r AliasRule) Tags() string {
	return r.tags
}

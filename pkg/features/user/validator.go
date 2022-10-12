package user

import (
	"context"

	"github.com/wascript3r/reservio/pkg/validator"
)

type Validator interface {
	RawRequest(s any) error
	EmailUniqueness(ctx context.Context, email string) error
}

type SharedValidator interface {
	GetRules() []validator.Rule
	Validator
}

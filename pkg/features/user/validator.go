package user

import "context"

type Validator interface {
	RawRequest(s any) error
	EmailUniqueness(ctx context.Context, email string) error
}

package repository

import "context"

type Transactor interface {
	WithinTx(context.Context, func(ctx context.Context) error) error
}

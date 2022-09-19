package repository

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/company/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert = "INSERT INTO companies (user_id, name, address, description) VALUES ($1, $2, $3, $4)"
)

type PgRepo struct {
	db *pgsql.Database
}

func NewPgRepo(db *pgsql.Database) *PgRepo {
	return &PgRepo{db}
}

func (p *PgRepo) Insert(ctx context.Context, cs *models.Company) error {
	_, err := p.db.ExecContext(
		ctx,
		insert,

		cs.UserID,
		cs.Name,
		cs.Address,
		cs.Description,
	)
	return pgsql.ParseWriteErr(err)
}

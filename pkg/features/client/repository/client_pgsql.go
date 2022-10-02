package repository

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/client/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert = "INSERT INTO clients (id, first_name, last_name, phone) VALUES ($1, $2, $3, $4)"
	get    = "SELECT id, first_name, last_name, phone FROM clients WHERE id = $1"
)

type PgRepo struct {
	db *pgsql.Database
}

func NewPgRepo(db *pgsql.Database) *PgRepo {
	return &PgRepo{db}
}

func (p *PgRepo) Insert(ctx context.Context, cs *models.Client) error {
	_, err := p.db.ExecContext(
		ctx,
		insert,

		cs.ID,
		cs.FirstName,
		cs.LastName,
		cs.Phone,
	)
	return pgsql.ParseWriteErr(err)
}

func (p *PgRepo) Get(ctx context.Context, id string) (*models.Client, error) {
	cs := &models.Client{}
	err := p.db.QueryRowContext(ctx, get, id).Scan(
		&cs.ID,
		&cs.FirstName,
		&cs.LastName,
		&cs.Phone,
	)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	return cs, nil
}

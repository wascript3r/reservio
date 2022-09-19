package repository

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/user/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert      = "INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id"
	emailExists = "SELECT EXISTS(SELECT 1 FROM users WHERE LOWER(email) = LOWER($1))"
)

type PgRepo struct {
	db *pgsql.Database
}

func NewPgRepo(db *pgsql.Database) *PgRepo {
	return &PgRepo{db}
}

func (p *PgRepo) Insert(ctx context.Context, us *models.User) (string, error) {
	var id string
	err := p.db.QueryRowContext(
		ctx,
		insert,

		us.Email,
		us.Password,
		us.Role,
	).Scan(&id)

	return id, pgsql.ParseWriteErr(err)
}

func (p *PgRepo) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := p.db.QueryRowContext(ctx, emailExists, email).Scan(&exists)
	return exists, err
}

package repository

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/user/models"
	"github.com/wascript3r/reservio/pkg/repository"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert      = "INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id"
	emailExists = "SELECT EXISTS(SELECT 1 FROM users WHERE LOWER(email) = LOWER($1))"
	getByEmail  = "SELECT id, email, password, role FROM users WHERE LOWER(email) = LOWER($1)"
	deleteq     = "DELETE FROM users WHERE id = $1"
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

func (p *PgRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	us := &models.User{}
	err := p.db.QueryRowContext(ctx, getByEmail, email).Scan(
		&us.ID,
		&us.Email,
		&us.Password,
		&us.Role,
	)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	return us, nil
}

func (p *PgRepo) Delete(ctx context.Context, id string) error {
	res, err := p.db.ExecContext(ctx, deleteq, id)
	if err != nil {
		return pgsql.ParseWriteErr(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	} else if n == 0 {
		return repository.ErrNoItems
	}

	return nil
}

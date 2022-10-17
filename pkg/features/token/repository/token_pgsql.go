package repository

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/token/models"
	"github.com/wascript3r/reservio/pkg/repository"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert         = "INSERT INTO refresh_tokens (user_id, expires_at) VALUES ($1, $2) RETURNING id"
	getClaims      = "SELECT rt.user_id, u.role FROM refresh_tokens rt JOIN users u ON u.id = rt.user_id WHERE rt.id = $1 AND rt.expires_at > $2"
	deleteq        = "DELETE FROM refresh_tokens WHERE id = $1"
	deleteByUserID = "DELETE FROM refresh_tokens WHERE user_id = $1"
)

type PgRepo struct {
	db *pgsql.Database
}

func NewPgRepo(db *pgsql.Database) *PgRepo {
	return &PgRepo{db}
}

func (p *PgRepo) Insert(ctx context.Context, rts *models.RefreshToken) (string, error) {
	var id string
	err := p.db.QueryRowContext(ctx, insert, rts.UserID, rts.ExpiresAt).Scan(&id)
	return id, pgsql.ParseWriteErr(err)
}

func (p *PgRepo) GetClaims(ctx context.Context, refreshTokenID string) (*models.Claims, error) {
	cs := &models.Claims{}
	err := p.db.QueryRowContext(ctx, getClaims, refreshTokenID, time.Now()).Scan(&cs.UserID, &cs.Role)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	return cs, nil
}

func (p *PgRepo) Delete(ctx context.Context, refreshTokenID string) error {
	res, err := p.db.ExecContext(ctx, deleteq, refreshTokenID)
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

func (p *PgRepo) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := p.db.ExecContext(ctx, deleteByUserID, userID)
	return pgsql.ParseWriteErr(err)
}

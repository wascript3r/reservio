package repository

import (
	"context"

	"github.com/wascript3r/reservio/pkg/repository"

	"github.com/wascript3r/reservio/pkg/features/company/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert = "INSERT INTO companies (user_id, name, address, description) VALUES ($1, $2, $3, $4)"
	get    = "SELECT c.user_id, c.name, c.address, c.description, u.email FROM companies c INNER JOIN users u ON u.id = c.user_id WHERE c.user_id = $1"
	getAll = "SELECT c.user_id, c.name, c.address, c.description, u.email FROM companies c INNER JOIN users u ON u.id = c.user_id ORDER BY c.created_at DESC"
	update = "UPDATE companies SET name = $2, address = $3, description = $4 WHERE user_id = $1"
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

func scanInfo(row pgsql.Row) (*models.CompanyInfo, error) {
	ci := &models.CompanyInfo{}
	err := row.Scan(
		&ci.Email,
		&ci.UserID,
		&ci.Name,
		&ci.Address,
		&ci.Description,
	)

	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	return ci, nil
}

func (p *PgRepo) Get(ctx context.Context, userID string) (*models.CompanyInfo, error) {
	row := p.db.QueryRowContext(ctx, get, userID)
	return scanInfo(row)
}

func (p *PgRepo) GetAll(ctx context.Context) ([]*models.CompanyInfo, error) {
	rows, err := p.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	defer rows.Close()

	var cis []*models.CompanyInfo
	for rows.Next() {
		ci, err := scanInfo(rows)
		if err != nil {
			return nil, err
		}
		cis = append(cis, ci)
	}

	return cis, nil
}

func (p *PgRepo) Update(ctx context.Context, userID string, cu *models.CompanyUpdate) error {
	res, err := p.db.ExecContext(
		ctx,
		update,

		userID,
		cu.Name,
		cu.Address,
		cu.Description,
	)
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

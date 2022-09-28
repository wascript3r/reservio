package repository

import (
	"context"

	"github.com/wascript3r/reservio/pkg/features/company/models"
	"github.com/wascript3r/reservio/pkg/repository"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert         = "INSERT INTO companies (user_id, name, address, description) VALUES ($1, $2, $3, $4)"
	get            = "SELECT u.email, c.user_id, c.name, c.address, c.description, c.approved FROM companies c INNER JOIN users u ON u.id = c.user_id WHERE c.user_id = $1"
	getApproved    = get + " AND c.approved = true"
	getAll         = "SELECT u.email, c.user_id, c.name, c.address, c.description, c.approved FROM companies c INNER JOIN users u ON u.id = c.user_id ORDER BY c.created_at DESC"
	getAllApproved = "SELECT u.email, c.user_id, c.name, c.address, c.description, c.approved FROM companies c INNER JOIN users u ON u.id = c.user_id WHERE c.approved = TRUE ORDER BY c.created_at DESC"

	update         = "UPDATE companies <set> WHERE user_id = $1"
	setName        = "name = ?"
	setAddress     = "address = ?"
	setDescription = "description = ?"

	deleteq = "DELETE FROM companies WHERE user_id = $1"
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
		&ci.Approved,
	)

	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	return ci, nil
}

func (p *PgRepo) Get(ctx context.Context, userID string, onlyApproved bool) (*models.CompanyInfo, error) {
	q := get
	if onlyApproved {
		q = getApproved
	}

	row := p.db.QueryRowContext(ctx, q, userID)
	return scanInfo(row)
}

func (p *PgRepo) GetAll(ctx context.Context, onlyApproved bool) ([]*models.CompanyInfo, error) {
	q := getAll
	if onlyApproved {
		q = getAllApproved
	}

	rows, err := p.db.QueryContext(ctx, q)
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
	if cu.IsEmpty() {
		return repository.ErrInvalidParamInput
	}
	builder := pgsql.NewQueryBuilder(pgsql.UpdateQuery, update, 2)

	if cu.Name != nil {
		builder.Add(setName, *cu.Name)
	}
	if cu.Address != nil {
		builder.Add(setAddress, *cu.Address)
	}
	if cu.Description != nil {
		builder.Add(setDescription, *cu.Description)
	}

	res, err := p.db.ExecContext(ctx, builder.GetQuery(), builder.GetParams(userID)...)
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

func (p *PgRepo) Delete(ctx context.Context, userID string) error {
	res, err := p.db.ExecContext(ctx, deleteq, userID)
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

package repository

import (
	"context"
	"encoding/json"

	"github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert = "INSERT INTO services (company_id, title, description, specialist_name, specialist_phone, work_schedule) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
)

type PgRepo struct {
	db *pgsql.Database
}

func NewPgRepo(db *pgsql.Database) *PgRepo {
	return &PgRepo{db}
}

func (p *PgRepo) Insert(ctx context.Context, ss *models.Service) (string, error) {
	bs, err := json.Marshal(ss.WorkSchedule)
	if err != nil {
		return "", err
	}

	var id string
	err = p.db.QueryRowContext(
		ctx,
		insert,

		ss.CompanyID,
		ss.Title,
		ss.Description,
		ss.SpecialistName,
		ss.SpecialistPhone,
		bs,
	).Scan(&id)

	return id, pgsql.ParseWriteErr(err)
}

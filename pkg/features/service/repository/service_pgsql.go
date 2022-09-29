package repository

import (
	"context"
	"encoding/json"

	"github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert      = "INSERT INTO services (company_id, title, description, specialist_name, specialist_phone, work_schedule) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	get         = "SELECT id, company_id, title, description, specialist_name, specialist_phone, work_schedule FROM services WHERE company_id = $1 AND id = $2"
	getApproved = "SELECT s.id, s.company_id, s.title, s.description, s.specialist_name, s.specialist_phone, s.work_schedule FROM services s INNER JOIN companies c ON c.id = s.company_id WHERE s.company_id = $1 AND s.id = $2 AND c.approved = TRUE"
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

func scanService(row pgsql.Row) (*models.Service, error) {
	s := &models.Service{}
	var bs []byte
	err := row.Scan(
		&s.ID,
		&s.CompanyID,
		&s.Title,
		&s.Description,
		&s.SpecialistName,
		&s.SpecialistPhone,
		&bs,
	)

	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}

	err = json.Unmarshal(bs, &s.WorkSchedule)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (p *PgRepo) Get(ctx context.Context, companyID, serviceID string, onlyApprovedCompany bool) (*models.Service, error) {
	q := get
	if onlyApprovedCompany {
		q = getApproved
	}

	row := p.db.QueryRowContext(ctx, q, companyID, serviceID)
	return scanService(row)
}

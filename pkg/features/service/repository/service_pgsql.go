package repository

import (
	"context"
	"encoding/json"

	"github.com/wascript3r/reservio/pkg/repository"

	"github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert         = "INSERT INTO services (company_id, title, description, specialist_name, specialist_phone, work_schedule) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	get            = "SELECT id, company_id, title, description, specialist_name, specialist_phone, work_schedule FROM services WHERE company_id = $1 AND id = $2"
	getApproved    = "SELECT s.id, s.company_id, s.title, s.description, s.specialist_name, s.specialist_phone, s.work_schedule FROM services s INNER JOIN companies c ON c.id = s.company_id WHERE s.company_id = $1 AND s.id = $2 AND c.approved = TRUE"
	getAll         = "SELECT id, company_id, title, description, specialist_name, specialist_phone, work_schedule FROM services WHERE company_id = $1 ORDER BY created_at DESC"
	getAllApproved = "SELECT s.id, s.company_id, s.title, s.description, s.specialist_name, s.specialist_phone, s.work_schedule FROM services s INNER JOIN companies c ON c.id = s.company_id WHERE s.company_id = $1 AND c.approved = TRUE ORDER BY s.created_at DESC"

	update             = "UPDATE services <set> WHERE company_id = $1 AND id = $2"
	setTitle           = "title = ?"
	setDescription     = "description = ?"
	setSpecialistName  = "specialist_name = ?"
	setSpecialistPhone = "specialist_phone = ?"
	setWorkSchedule    = "work_schedule = ?"

	deleteq           = "DELETE FROM services WHERE company_id = $1 AND id = $2"
	deleteByCompanyID = "DELETE FROM services WHERE company_id = $1"
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

func (p *PgRepo) GetAll(ctx context.Context, companyID string, onlyApprovedCompany bool) ([]*models.Service, error) {
	q := getAll
	if onlyApprovedCompany {
		q = getAllApproved
	}

	rows, err := p.db.QueryContext(ctx, q, companyID)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	defer rows.Close()

	var services []*models.Service
	for rows.Next() {
		s, err := scanService(rows)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}

	return services, nil
}

func (p *PgRepo) Update(ctx context.Context, companyID, serviceID string, su *models.ServiceUpdate) error {
	if su.IsEmpty() {
		return repository.ErrInvalidParamInput
	}
	builder := pgsql.NewQueryBuilder(pgsql.UpdateQuery, update, 3)

	if su.Title != nil {
		builder.Add(setTitle, *su.Title)
	}
	if su.Description != nil {
		builder.Add(setDescription, *su.Description)
	}
	if su.SpecialistName != nil {
		builder.Add(setSpecialistName, *su.SpecialistName)
	}
	if su.SpecialistPhone != nil {
		builder.Add(setSpecialistPhone, *su.SpecialistPhone)
	}
	if su.WorkSchedule != nil {
		bs, err := json.Marshal(su.WorkSchedule)
		if err != nil {
			return err
		}
		builder.Add(setWorkSchedule, bs)
	}

	res, err := p.db.ExecContext(ctx, builder.GetQuery(), builder.GetParams(companyID, serviceID)...)
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

func (p *PgRepo) Delete(ctx context.Context, companyID, serviceID string) error {
	res, err := p.db.ExecContext(ctx, deleteq, companyID, serviceID)
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

func (p *PgRepo) DeleteByCompany(ctx context.Context, companyID string) error {
	_, err := p.db.ExecContext(ctx, deleteByCompanyID, companyID)
	return pgsql.ParseWriteErr(err)
}

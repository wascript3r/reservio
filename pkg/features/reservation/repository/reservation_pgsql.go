package repository

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/repository"

	"github.com/wascript3r/reservio/pkg/features/reservation/models"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert         = "INSERT INTO reservations (service_id, date, comment) VALUES ($1, $2, $3) RETURNING id"
	get            = "SELECT r.id, r.service_id, r.date, r.comment FROM reservations r INNER JOIN services s ON s.id = r.service_id WHERE s.company_id = $1 AND s.id = $2 AND r.id = $3"
	getApproved    = "SELECT r.id, r.service_id, r.date, r.comment FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN companies c ON c.id = s.company_id WHERE s.company_id = $1 AND s.id = $2 AND r.id = $3 AND c.approved = TRUE"
	getAll         = "SELECT r.id, r.service_id, r.date, r.comment FROM reservations r INNER JOIN services s ON s.id = r.service_id WHERE s.company_id = $1 AND s.id = $2 ORDER BY r.date"
	getAllApproved = "SELECT r.id, r.service_id, r.date, r.comment FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN companies c ON c.id = s.company_id WHERE s.company_id = $1 AND s.id = $2 AND c.approved = TRUE ORDER BY r.date"

	update     = "UPDATE reservations r <set> FROM services s WHERE s.id = r.service_id AND s.company_id = $1 AND s.id = $2 AND r.id = $3"
	setDate    = "date = ?"
	setComment = "comment = ?"

	deleteq = "DELETE FROM reservations r USING services s WHERE s.id = r.service_id AND s.company_id = $1 AND s.id = $2 AND r.id = $3"
	exists  = "SELECT EXISTS (SELECT 1 FROM reservations r INNER JOIN services s ON s.id = r.service_id WHERE s.company_id = $1 AND s.id = $2 AND r.date = $3)"
)

type PgRepo struct {
	db *pgsql.Database
}

func NewPgRepo(db *pgsql.Database) *PgRepo {
	return &PgRepo{db}
}

func (p *PgRepo) Insert(ctx context.Context, rs *models.Reservation) (string, error) {
	var id string
	err := p.db.QueryRowContext(
		ctx,
		insert,

		rs.ServiceID,
		rs.Date,
		rs.Comment,
	).Scan(&id)

	return id, pgsql.ParseWriteErr(err)
}

func scanReservation(row pgsql.Row) (*models.Reservation, error) {
	r := &models.Reservation{}
	err := row.Scan(&r.ID, &r.ServiceID, &r.Date, &r.Comment)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	return r, nil
}

func (p *PgRepo) Get(ctx context.Context, companyID, serviceID, reservationID string, onlyApprovedCompany bool) (*models.Reservation, error) {
	q := get
	if onlyApprovedCompany {
		q = getApproved
	}

	row := p.db.QueryRowContext(ctx, q, companyID, serviceID, reservationID)
	return scanReservation(row)
}

func (p *PgRepo) GetAll(ctx context.Context, companyID, serviceID string, onlyApprovedCompany bool) ([]*models.Reservation, error) {
	q := getAll
	if onlyApprovedCompany {
		q = getAllApproved
	}

	rows, err := p.db.QueryContext(ctx, q, companyID, serviceID)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	defer rows.Close()

	var reservations []*models.Reservation
	for rows.Next() {
		r, err := scanReservation(rows)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, r)
	}

	return reservations, nil
}

func (p *PgRepo) Update(ctx context.Context, companyID, serviceID, reservationID string, ru *models.ReservationUpdate) error {
	if ru.IsEmpty() {
		return repository.ErrInvalidParamInput
	}
	builder := pgsql.NewQueryBuilder(pgsql.UpdateQuery, update, 4)

	if ru.Date != nil {
		builder.Add(setDate, *ru.Date)
	}
	if ru.Comment != nil {
		builder.Add(setComment, *ru.Comment)
	}

	res, err := p.db.ExecContext(ctx, builder.GetQuery(), builder.GetParams(companyID, serviceID, reservationID)...)
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

func (p *PgRepo) Delete(ctx context.Context, companyID, serviceID, reservationID string) error {
	res, err := p.db.ExecContext(ctx, deleteq, companyID, serviceID, reservationID)
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

func (p *PgRepo) Exists(ctx context.Context, companyID, serviceID string, date time.Time) (bool, error) {
	var ex bool
	err := p.db.QueryRowContext(ctx, exists, companyID, serviceID, date).Scan(&ex)
	if err != nil {
		return false, pgsql.ParseReadErr(err)
	}
	return ex, nil
}

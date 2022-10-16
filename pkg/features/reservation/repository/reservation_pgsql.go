package repository

import (
	"context"
	"encoding/json"
	"time"

	clmodels "github.com/wascript3r/reservio/pkg/features/client/models"
	cmodels "github.com/wascript3r/reservio/pkg/features/company/models"
	"github.com/wascript3r/reservio/pkg/features/reservation/models"
	smodels "github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/repository"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const (
	insert = "INSERT INTO reservations (service_id, client_id, date, comment) VALUES ($1, $2, $3, $4) RETURNING id"

	get         = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 AND s.id = $2 AND r.id = $3"
	getByClient = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 AND s.id = $2 AND r.id = $3 AND r.client_id = $4"

	getApproved         = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN companies c ON c.id = s.company_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 AND s.id = $2 AND r.id = $3 AND c.approved = TRUE"
	getApprovedByClient = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN companies c ON c.id = s.company_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 AND s.id = $2 AND r.id = $3 AND c.approved = TRUE AND r.client_id = $4"

	getAll                  = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 AND s.id = $2 ORDER BY r.date"
	getAllApproved          = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN companies c ON c.id = s.company_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 AND s.id = $2 AND c.approved = TRUE ORDER BY r.date"
	getAllByCompany         = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 ORDER BY r.date"
	getAllByCompanyApproved = "SELECT r.id, r.service_id, r.date, r.comment, cl.id, cl.first_name, cl.last_name, cl.phone, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN companies c ON c.id = s.company_id INNER JOIN clients cl ON cl.id = r.client_id INNER JOIN users u ON u.id = cl.id WHERE s.company_id = $1 AND c.approved = TRUE ORDER BY r.date"

	getAllByClient = "SELECT r.id, r.date, r.comment, s.id, s.title, s.description, s.specialist_name, s.specialist_phone, s.visit_duration, s.work_schedule, c.id, c.name, c.address, c.description, c.approved, u.email FROM reservations r INNER JOIN services s ON s.id = r.service_id INNER JOIN companies c ON c.id = s.company_id INNER JOIN users u ON u.id = c.id WHERE r.client_id = $1 ORDER BY r.date"

	update     = "UPDATE reservations r <set> FROM services s WHERE s.id = r.service_id AND s.company_id = $1 AND s.id = $2 AND r.id = $3 AND r.client_id = $4"
	setDate    = "date = ?"
	setComment = "comment = ?"

	deleteq         = "DELETE FROM reservations r USING services s WHERE s.id = r.service_id AND s.company_id = $1 AND s.id = $2 AND r.id = $3 AND r.client_id = $4"
	exists          = "SELECT EXISTS (SELECT 1 FROM reservations r INNER JOIN services s ON s.id = r.service_id WHERE s.company_id = $1 AND s.id = $2 AND r.date = $3)"
	deleteByCompany = "DELETE FROM reservations r USING services s WHERE s.id = r.service_id AND s.company_id = $1"
	deleteByService = "DELETE FROM reservations WHERE service_id = $1"
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
		rs.ClientID,
		rs.Date,
		rs.Comment,
	).Scan(&id)

	return id, pgsql.ParseWriteErr(err)
}

func scanReservation(row pgsql.Row) (*models.FullReservation, error) {
	r := &models.FullReservation{
		Client: &clmodels.ClientInfo{},
	}
	err := row.Scan(
		&r.ID,
		&r.ServiceID,
		&r.Date,
		&r.Comment,
		&r.Client.ID,
		&r.Client.FirstName,
		&r.Client.LastName,
		&r.Client.Phone,
		&r.Client.Email,
	)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	return r, nil
}

func (p *PgRepo) Get(ctx context.Context, companyID, serviceID, reservationID string, clientID *string, onlyApprovedCompany bool) (*models.FullReservation, error) {
	var (
		q    string
		args = []interface{}{companyID, serviceID, reservationID}
	)
	if clientID != nil {
		if onlyApprovedCompany {
			q = getApprovedByClient
		} else {
			q = getByClient
		}
		args = append(args, *clientID)
	} else {
		if onlyApprovedCompany {
			q = getApproved
		} else {
			q = get
		}
	}

	row := p.db.QueryRowContext(ctx, q, args...)
	return scanReservation(row)
}

func (p *PgRepo) GetAll(ctx context.Context, companyID, serviceID string, onlyApprovedCompany bool) ([]*models.FullReservation, error) {
	q := getAll
	if onlyApprovedCompany {
		q = getAllApproved
	}

	rows, err := p.db.QueryContext(ctx, q, companyID, serviceID)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	defer rows.Close()

	var reservations []*models.FullReservation
	for rows.Next() {
		r, err := scanReservation(rows)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, r)
	}

	return reservations, nil
}

func (p *PgRepo) GetAllByCompany(ctx context.Context, companyID string, onlyApprovedCompany bool) ([]*models.FullReservation, error) {
	q := getAllByCompany
	if onlyApprovedCompany {
		q = getAllByCompanyApproved
	}

	rows, err := p.db.QueryContext(ctx, q, companyID)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	defer rows.Close()

	var reservations []*models.FullReservation
	for rows.Next() {
		r, err := scanReservation(rows)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, r)
	}

	return reservations, nil
}

func (p *PgRepo) GetAllByClient(ctx context.Context, clientID string) ([]*models.ClientReservation, error) {
	rows, err := p.db.QueryContext(ctx, getAllByClient, clientID)
	if err != nil {
		return nil, pgsql.ParseReadErr(err)
	}
	defer rows.Close()

	var reservations []*models.ClientReservation
	for rows.Next() {
		var bs []byte
		r := &models.ClientReservation{
			Service: &smodels.FullService{
				Company: &cmodels.CompanyInfo{},
			},
		}
		err := rows.Scan(
			&r.ID,
			&r.Date,
			&r.Comment,
			&r.Service.ID,
			&r.Service.Title,
			&r.Service.Description,
			&r.Service.SpecialistName,
			&r.Service.SpecialistPhone,
			&r.Service.VisitDuration,
			&bs,
			&r.Service.Company.ID,
			&r.Service.Company.Name,
			&r.Service.Company.Address,
			&r.Service.Company.Description,
			&r.Service.Company.Approved,
			&r.Service.Company.Email,
		)
		if err != nil {
			return nil, pgsql.ParseReadErr(err)
		}
		err = json.Unmarshal(bs, &r.Service.WorkSchedule)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, r)
	}

	return reservations, nil
}

func (p *PgRepo) Update(ctx context.Context, companyID, serviceID, reservationID, clientID string, ru *models.ReservationUpdate) error {
	if ru.IsEmpty() {
		return repository.ErrInvalidParamInput
	}
	builder := pgsql.NewQueryBuilder(pgsql.UpdateQuery, update, 5)

	if ru.Date != nil {
		builder.Add(setDate, *ru.Date)
	}
	if ru.Comment != nil {
		builder.Add(setComment, *ru.Comment)
	}

	res, err := p.db.ExecContext(ctx, builder.GetQuery(), builder.GetParams(companyID, serviceID, reservationID, clientID)...)
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

func (p *PgRepo) Delete(ctx context.Context, companyID, serviceID, reservationID, clientID string) error {
	res, err := p.db.ExecContext(ctx, deleteq, companyID, serviceID, reservationID, clientID)
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

func (p *PgRepo) DeleteByCompany(ctx context.Context, companyID string) error {
	_, err := p.db.ExecContext(ctx, deleteByCompany, companyID)
	return pgsql.ParseWriteErr(err)
}

func (p *PgRepo) DeleteByService(ctx context.Context, serviceID string) error {
	_, err := p.db.ExecContext(ctx, deleteByService, serviceID)
	return pgsql.ParseWriteErr(err)
}

package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	cmodels "github.com/wascript3r/reservio/pkg/features/client/models"
	rmodels "github.com/wascript3r/reservio/pkg/features/reservation/models"
)

const (
	companyID = "e91f4c92-1371-48a1-a745-7d66d2178e15"
	serviceID = "1b618d1e-05cc-4d10-84fe-c304d424afa0"
)

func newReservation() *rmodels.FullReservation {
	comment := "Man kad švaru būtų!!!"
	return &rmodels.FullReservation{
		ID:        "136c80c7-7d4c-4edf-a7ea-e06ffc25505e",
		ServiceID: "1b618d1e-05cc-4d10-84fe-c304d424afa0",
		Client: &cmodels.ClientInfo{
			Client: cmodels.Client{
				ID:        "d974f2b6-dfa2-4d03-98fd-d318c4262e68",
				FirstName: "Petras",
				LastName:  "Petraitis",
				Phone:     "+37061354544",
			},
			Email: "petras.petraitis@gmail.com",
		},
		Date:    time.Date(2022, 10, 10, 8, 0, 0, 0, time.UTC),
		Comment: &comment,
	}
}

type ReservationRepoSuite struct {
	suite.Suite

	db   *sql.DB
	mock sqlmock.Sqlmock
	repo *PgRepo
}

func (r *ReservationRepoSuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	require.NoError(r.T(), err)

	r.db = db
	r.mock = mock
	r.repo = NewPgRepo(db)
}

func (r *ReservationRepoSuite) TestGetAllApproved() {
	rs := newReservation()

	// language=ignorelang
	q := "SELECT (.+) FROM reservations (.+) WHERE (.*)approved = TRUE"
	rows := sqlmock.NewRows([]string{
		"id", "service_id", "date", "comment", "client_id", "first_name", "last_name", "phone", "email",
	}).AddRow(rs.ID, rs.ServiceID, rs.Date, rs.Comment, rs.Client.ID, rs.Client.FirstName, rs.Client.LastName, rs.Client.Phone, rs.Client.Email)

	r.mock.
		ExpectQuery(q).
		WithArgs(companyID, serviceID).
		WillReturnRows(rows).
		RowsWillBeClosed()

	rss, err := r.repo.GetAll(context.Background(), companyID, serviceID, true)

	require.NoError(r.T(), err)
	require.Len(r.T(), rss, 1)
	require.Equal(r.T(), rs, rss[0])
}

func TestReservationRepo(t *testing.T) {
	suite.Run(t, &ReservationRepoSuite{})
}

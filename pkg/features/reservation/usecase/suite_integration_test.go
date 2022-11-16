package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/wascript3r/reservio/cmd/app/config"
	clrepo "github.com/wascript3r/reservio/pkg/features/client/repository"
	"github.com/wascript3r/reservio/pkg/features/company/dto"
	crepo "github.com/wascript3r/reservio/pkg/features/company/repository"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	rdto "github.com/wascript3r/reservio/pkg/features/reservation/dto"
	rrepo "github.com/wascript3r/reservio/pkg/features/reservation/repository"
	"github.com/wascript3r/reservio/pkg/features/reservation/validator"
	sdto "github.com/wascript3r/reservio/pkg/features/service/dto"
	srepo "github.com/wascript3r/reservio/pkg/features/service/repository"
	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

const dbDriver = "postgres"

type UcaseIntegrationSuite struct {
	suite.Suite
	cfg    *config.Config
	dbConn *sql.DB

	reservationRepo *rrepo.PgRepo
	serviceRepo     *srepo.PgRepo
	companyRepo     *crepo.PgRepo
	clientRepo      *clrepo.PgRepo
	ctxTimeout      time.Duration
	validator       *validator.Validator

	ctx     context.Context
	usecase *Usecase
}

func (u *UcaseIntegrationSuite) SetupSuite() {
	cfg, err := config.LoadConfig(true)
	require.NoError(u.T(), err)

	u.cfg = cfg
	u.ctxTimeout = cfg.Database.Postgres.QueryTimeout.Duration

	txdb.Register("txdb", dbDriver, u.cfg.Database.Postgres.Integration.DSN)
}

func (u *UcaseIntegrationSuite) SetupTest() {
	conn, err := sql.Open("txdb", uuid.NewString())
	require.NoError(u.T(), err)
	require.NoError(u.T(), conn.Ping())

	u.dbConn = conn
	pgdb := pgsql.NewDatabase(conn)

	u.reservationRepo = rrepo.NewPgRepo(pgdb)
	u.serviceRepo = srepo.NewPgRepo(pgdb)
	u.companyRepo = crepo.NewPgRepo(pgdb)
	u.clientRepo = clrepo.NewPgRepo(pgdb)
	u.validator = validator.New()

	u.ctx = context.Background()
	u.usecase = New(u.reservationRepo, u.serviceRepo, u.companyRepo, u.clientRepo, u.ctxTimeout, u.validator)
}

func (u *UcaseIntegrationSuite) TearDownTest() {
	require.NoError(u.T(), u.dbConn.Close())
}

func (u *UcaseIntegrationSuite) TestUpdate() {
	date := "2030-01-02 10:00"
	upd := &rdto.UpdateReq{
		ReservationReq: rdto.ReservationReq{
			ServiceReq: sdto.ServiceReq{
				CompanyReq: dto.CompanyReq{CompanyID: companyID},
				ServiceID:  serviceID,
			},
			ReservationID: reservationID,
		},
		ClientID: clientID,
		Date:     &date,
		Comment:  &rdto.Comment{Value: nil},
	}
	get := &rdto.GetReq{
		ReservationReq: upd.ReservationReq,
		ClientID:       &upd.ClientID,
	}

	res, err := u.usecase.Get(u.ctx, get, false)
	require.NoError(u.T(), err)
	require.NotEqual(u.T(), *upd.Date, res.Date)
	require.NotEqual(u.T(), upd.Comment, res.Comment)

	err = u.usecase.Update(u.ctx, upd)
	require.NoError(u.T(), err)

	res, err = u.usecase.Get(u.ctx, get, false)
	require.NoError(u.T(), err)
	require.Equal(u.T(), *upd.Date, res.Date)
	require.Equal(u.T(), upd.Comment.Value, res.Comment)
}

func (u *UcaseIntegrationSuite) TestDelete() {
	del := &rdto.DeleteReq{
		ReservationReq: rdto.ReservationReq{
			ServiceReq: sdto.ServiceReq{
				CompanyReq: dto.CompanyReq{CompanyID: companyID},
				ServiceID:  serviceID,
			},
			ReservationID: reservationID,
		},
		ClientID: clientID,
	}
	get := &rdto.GetReq{
		ReservationReq: del.ReservationReq,
		ClientID:       &del.ClientID,
	}

	res, err := u.usecase.Get(u.ctx, get, false)
	require.NoError(u.T(), err)
	require.NotNil(u.T(), res)

	err = u.usecase.Delete(u.ctx, del)
	require.NoError(u.T(), err)

	res, err = u.usecase.Get(u.ctx, get, false)
	require.ErrorIs(u.T(), err, reservation.NotFoundError)
	require.Nil(u.T(), res)

	err = u.usecase.Delete(u.ctx, del)
	require.ErrorIs(u.T(), err, reservation.NotFoundError)
}

func TestReservationUcaseIntegration(t *testing.T) {
	suite.Run(t, &UcaseIntegrationSuite{})
}

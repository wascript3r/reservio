package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	clmocks "github.com/wascript3r/reservio/pkg/features/client/mocks"
	cmocks "github.com/wascript3r/reservio/pkg/features/company/mocks"
	rmocks "github.com/wascript3r/reservio/pkg/features/reservation/mocks"
	smocks "github.com/wascript3r/reservio/pkg/features/service/mocks"
)

const (
	companyID     = "e91f4c92-1371-48a1-a745-7d66d2178e15"
	serviceID     = "1b618d1e-05cc-4d10-84fe-c304d424afa0"
	clientID      = "d974f2b6-dfa2-4d03-98fd-d318c4262e68"
	reservationID = "136c80c7-7d4c-4edf-a7ea-e06ffc25505e"
)

type UcaseSuite struct {
	suite.Suite

	reservationRepo *rmocks.MockRepository
	serviceRepo     *smocks.MockRepository
	companyRepo     *cmocks.MockRepository
	clientRepo      *clmocks.MockRepository
	ctxTimeout      time.Duration
	validator       *rmocks.MockValidator

	ctx     context.Context
	usecase *Usecase
}

func (u *UcaseSuite) SetupSuite() {
	u.ctxTimeout = 10 * time.Second
}

func (u *UcaseSuite) SetupTest() {
	ctrl := gomock.NewController(u.T())

	u.reservationRepo = rmocks.NewMockRepository(ctrl)
	u.serviceRepo = smocks.NewMockRepository(ctrl)
	u.companyRepo = cmocks.NewMockRepository(ctrl)
	u.clientRepo = clmocks.NewMockRepository(ctrl)
	u.validator = rmocks.NewMockValidator(ctrl)

	u.ctx = context.Background()
	u.usecase = New(u.reservationRepo, u.serviceRepo, u.companyRepo, u.clientRepo, u.ctxTimeout, u.validator)
}

func TestReservationUcase(t *testing.T) {
	suite.Run(t, &UcaseSuite{})
}

package usecase

import (
	"html"
	"time"

	"github.com/wascript3r/reservio/pkg/features/client"

	rmodels "github.com/wascript3r/reservio/pkg/features/reservation/models"

	mck "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/company/dto"
	cmodels "github.com/wascript3r/reservio/pkg/features/company/models"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	rdto "github.com/wascript3r/reservio/pkg/features/reservation/dto"
	"github.com/wascript3r/reservio/pkg/features/service"
	sdto "github.com/wascript3r/reservio/pkg/features/service/dto"
	smodels "github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/repository"
	"github.com/wascript3r/reservio/pkg/test"
)

func (r *ReservationUcaseSuite) TestCreate() {
	const validDate = "2022-01-01 11:11"
	comment := func() *string {
		c := "<h2>Test comment</h2>"
		return &c
	}

	cases := test.Cases[*rdto.CreateReq, *rdto.CreateRes]{
		{
			Name: "WrongDate",
			Prepare: func() {
				mck.InOrder(
					r.validator.EXPECT().RawRequest(mck.Any()).Return(nil),
				)
			},
			Req: &rdto.CreateReq{
				Date: "2022-01-01 11::11",
			},
			ExpectedErr: reservation.InvalidInputError,
			ValidateRes: nil,
		},
		{
			Name: "CompanyNotFound",
			Prepare: func() {
				mck.InOrder(
					r.validator.EXPECT().RawRequest(mck.Any()).Return(nil),
					r.companyRepo.EXPECT().Get(mck.Any(), companyID, false).
						Return(nil, repository.ErrNoItems),
				)
			},
			Req: &rdto.CreateReq{
				ServiceReq: sdto.ServiceReq{
					CompanyReq: dto.CompanyReq{CompanyID: companyID},
				},
				Date: validDate,
			},
			ExpectedErr: company.NotFoundError,
			ValidateRes: nil,
		},
		{
			Name: "CompanyNotApproved",
			Prepare: func() {
				mck.InOrder(
					r.validator.EXPECT().RawRequest(mck.Any()).Return(nil),
					r.companyRepo.EXPECT().Get(mck.Any(), companyID, false).Return(
						&cmodels.CompanyInfo{
							Approved: false,
						},
						nil,
					),
				)
			},
			Req: &rdto.CreateReq{
				ServiceReq: sdto.ServiceReq{
					CompanyReq: dto.CompanyReq{CompanyID: companyID},
				},
				Date: validDate,
			},
			ExpectedErr: company.NotApprovedError,
			ValidateRes: nil,
		},
		{
			Name: "ServiceNotFound",
			Prepare: func() {
				mck.InOrder(
					r.validator.EXPECT().RawRequest(mck.Any()).Return(nil),
					r.companyRepo.EXPECT().Get(mck.Any(), companyID, false).Return(
						&cmodels.CompanyInfo{
							Approved: true,
						},
						nil,
					),
					r.serviceRepo.EXPECT().Get(mck.Any(), companyID, serviceID, false).
						Return(nil, repository.ErrNoItems),
				)
			},
			Req: &rdto.CreateReq{
				ServiceReq: sdto.ServiceReq{
					CompanyReq: dto.CompanyReq{CompanyID: companyID},
					ServiceID:  serviceID,
				},
				Date: validDate,
			},
			ExpectedErr: service.NotFoundError,
			ValidateRes: nil,
		},
		{
			Name: "ReservationExists",
			Prepare: func() {
				ss := &smodels.Service{}
				parsedDate, err := time.Parse(dateFormat, validDate)
				require.NoError(r.T(), err)

				mck.InOrder(
					r.validator.EXPECT().RawRequest(mck.Any()).Return(nil),
					r.companyRepo.EXPECT().Get(mck.Any(), companyID, false).Return(
						&cmodels.CompanyInfo{
							Approved: true,
						},
						nil,
					),
					r.serviceRepo.EXPECT().Get(mck.Any(), companyID, serviceID, false).
						Return(ss, nil),
					r.validator.EXPECT().ReservationDate(ss, parsedDate).Return(nil),
					r.reservationRepo.EXPECT().Exists(mck.Any(), companyID, serviceID, parsedDate).
						Return(true, nil),
				)
			},
			Req: &rdto.CreateReq{
				ServiceReq: sdto.ServiceReq{
					CompanyReq: dto.CompanyReq{CompanyID: companyID},
					ServiceID:  serviceID,
				},
				Date: validDate,
			},
			ExpectedErr: reservation.AlreadyExistsError,
			ValidateRes: nil,
		},
		{
			Name: "ClientNotFound",
			Prepare: func() {
				ss := &smodels.Service{}
				parsedDate, err := time.Parse(dateFormat, validDate)
				require.NoError(r.T(), err)

				rs := &rmodels.Reservation{
					ServiceID: serviceID,
					ClientID:  clientID,
					Date:      parsedDate,
					Comment:   comment(),
				}
				*rs.Comment = html.EscapeString(*rs.Comment)

				mck.InOrder(
					r.validator.EXPECT().RawRequest(mck.Any()).Return(nil),
					r.companyRepo.EXPECT().Get(mck.Any(), companyID, false).Return(
						&cmodels.CompanyInfo{
							Approved: true,
						},
						nil,
					),
					r.serviceRepo.EXPECT().Get(mck.Any(), companyID, serviceID, false).
						Return(ss, nil),
					r.validator.EXPECT().ReservationDate(ss, parsedDate).Return(nil),
					r.reservationRepo.EXPECT().Exists(mck.Any(), companyID, serviceID, parsedDate).
						Return(false, nil),
					r.reservationRepo.EXPECT().Insert(mck.Any(), mck.Eq(rs)).
						Return("", repository.ErrConflictWithRelatedItems),
				)
			},
			Req: &rdto.CreateReq{
				ServiceReq: sdto.ServiceReq{
					CompanyReq: dto.CompanyReq{CompanyID: companyID},
					ServiceID:  serviceID,
				},
				ClientID: clientID,
				Date:     validDate,
				Comment:  comment(),
			},
			ExpectedErr: client.NotFoundError,
			ValidateRes: nil,
		},
		{
			Name: "Success",
			Prepare: func() {
				ss := &smodels.Service{}
				parsedDate, err := time.Parse(dateFormat, validDate)
				require.NoError(r.T(), err)

				rs := &rmodels.Reservation{
					ServiceID: serviceID,
					ClientID:  clientID,
					Date:      parsedDate,
					Comment:   comment(),
				}
				*rs.Comment = html.EscapeString(*rs.Comment)

				mck.InOrder(
					r.validator.EXPECT().RawRequest(mck.Any()).Return(nil),
					r.companyRepo.EXPECT().Get(mck.Any(), companyID, false).Return(
						&cmodels.CompanyInfo{
							Approved: true,
						},
						nil,
					),
					r.serviceRepo.EXPECT().Get(mck.Any(), companyID, serviceID, false).
						Return(ss, nil),
					r.validator.EXPECT().ReservationDate(ss, parsedDate).Return(nil),
					r.reservationRepo.EXPECT().Exists(mck.Any(), companyID, serviceID, parsedDate).
						Return(false, nil),
					r.reservationRepo.EXPECT().Insert(mck.Any(), mck.Eq(rs)).
						Return(reservationID, nil),
				)
			},
			Req: &rdto.CreateReq{
				ServiceReq: sdto.ServiceReq{
					CompanyReq: dto.CompanyReq{CompanyID: companyID},
					ServiceID:  serviceID,
				},
				ClientID: clientID,
				Date:     validDate,
				Comment:  comment(),
			},
			ExpectedErr: nil,
			ValidateRes: func(res *rdto.CreateRes) {
				exp := &rdto.CreateRes{
					ID:        reservationID,
					ServiceID: serviceID,
					ClientID:  clientID,
					Date:      validDate,
					Comment:   comment(),
				}
				*exp.Comment = html.EscapeString(*exp.Comment)
				require.Equal(r.T(), exp, res)
			},
		},
	}

	cases.Test(r, func(c *test.Case[*rdto.CreateReq, *rdto.CreateRes]) {
		res, err := r.usecase.Create(r.ctx, c.Req)
		c.Validate(r, res, err)
	})
}

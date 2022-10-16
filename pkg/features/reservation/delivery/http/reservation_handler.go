package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	"github.com/wascript3r/reservio/pkg/features/reservation/dto"
	"github.com/wascript3r/reservio/pkg/features/token"
	mid "github.com/wascript3r/reservio/pkg/features/token/delivery/http"
	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

const InitRoute = "/v1/companies/:companyID/services/:serviceID/reservations"

type HTTPHandler struct {
	mapper           *httpjson.CodeMapper
	reservationUcase reservation.Usecase
	tokenUcase       token.Usecase
}

func NewHTTPHandler(ctx context.Context, r *httprouter.Router, client mid.Client, companyOrClient mid.CompanyOrClient, parse mid.Parse, cm *httpjson.CodeMapper, ru reservation.Usecase, tu token.Usecase) {
	handler := &HTTPHandler{
		mapper:           cm,
		reservationUcase: ru,
		tokenUcase:       tu,
	}
	handler.initErrs()

	r.POST(InitRoute, client.Wrap(ctx, handler.Create))
	r.GET(InitRoute+"/:reservationID", companyOrClient.Wrap(ctx, handler.Get))
	r.GET(InitRoute, parse.Wrap(ctx, handler.GetAll))
	r.PATCH(InitRoute+"/:reservationID", client.Wrap(ctx, handler.Update))
	r.DELETE(InitRoute+"/:reservationID", client.Wrap(ctx, handler.Delete))
}

func (h *HTTPHandler) initErrs() {
	h.mapper.Register(
		http.StatusBadRequest,
		reservation.InvalidInputError,
		reservation.NothingToUpdateError,
		reservation.DateIsInPastError,
		reservation.InvalidTimeError,
		reservation.ServiceNotAvailableError,
	)
	h.mapper.Register(http.StatusNotFound, reservation.NotFoundError)
	h.mapper.Register(http.StatusUnprocessableEntity, reservation.AlreadyExistsError)
	h.mapper.Register(http.StatusInternalServerError, reservation.UnknownError)
}

func (h *HTTPHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.CreateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err != nil {
		httpjson.InternalError(w, nil)
		return
	}
	req.ClientID = claims.UserID

	res, err := h.reservationUcase.Create(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, reservation.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.GetReq{}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")
	req.ReservationID = p.ByName("reservationID")

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err != nil {
		httpjson.InternalError(w, nil)
		return
	} else if claims.Role == umodels.CompanyRole && claims.UserID != req.CompanyID {
		httpjson.Forbidden(w, nil)
		return
	} else if claims.Role == umodels.ClientRole {
		req.ClientID = &claims.UserID
	}

	res, err := h.reservationUcase.Get(r.Context(), req, false)
	if err != nil {
		code := errcode.UnwrapErr(err, reservation.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.GetAllReq{}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err == nil && claims.Role == umodels.CompanyRole && claims.UserID == req.CompanyID {
		res, err := h.reservationUcase.GetAll(r.Context(), req, false)
		if err != nil {
			code := errcode.UnwrapErr(err, reservation.UnknownError)
			h.mapper.ServeErr(w, code, nil)
			return
		}

		httpjson.ServeJSON(w, res)
	} else {
		res, err := h.reservationUcase.GetAllMeta(r.Context(), (*dto.GetAllMetaReq)(req), false)
		if err != nil {
			code := errcode.UnwrapErr(err, reservation.UnknownError)
			h.mapper.ServeErr(w, code, nil)
			return
		}

		httpjson.ServeJSON(w, res)
	}
}

func (h *HTTPHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.UpdateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")
	req.ReservationID = p.ByName("reservationID")

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err != nil {
		httpjson.InternalError(w, nil)
		return
	}
	req.ClientID = claims.UserID

	err = h.reservationUcase.Update(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, reservation.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

func (h *HTTPHandler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.DeleteReq{}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")
	req.ReservationID = p.ByName("reservationID")

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err != nil {
		httpjson.InternalError(w, nil)
		return
	}
	req.ClientID = claims.UserID

	err = h.reservationUcase.Delete(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, reservation.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

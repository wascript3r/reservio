package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
	"github.com/wascript3r/reservio/pkg/features/client"
	"github.com/wascript3r/reservio/pkg/features/client/dto"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	rdto "github.com/wascript3r/reservio/pkg/features/reservation/dto"
	"github.com/wascript3r/reservio/pkg/features/token"
	mid "github.com/wascript3r/reservio/pkg/features/token/delivery/http"
)

const InitRoute = "/v1/clients"

type HTTPHandler struct {
	mapper           *httpjson.CodeMapper
	clientUcase      client.Usecase
	reservationUcase reservation.Usecase
	tokenUcase       token.Usecase
}

func NewHTTPHandler(ctx context.Context, r *httprouter.Router, client mid.Client, cm *httpjson.CodeMapper, cu client.Usecase, ru reservation.Usecase, tu token.Usecase) {
	handler := &HTTPHandler{
		mapper:           cm,
		clientUcase:      cu,
		reservationUcase: ru,
		tokenUcase:       tu,
	}
	handler.initErrs()

	r.POST(InitRoute, handler.Create)
	r.GET(InitRoute+"/:id/reservations", client.Wrap(ctx, handler.GetReservations))
}

func (h *HTTPHandler) initErrs() {
	h.mapper.Register(http.StatusBadRequest, client.InvalidInputError)
	h.mapper.Register(http.StatusNotFound, client.NotFoundError)
	h.mapper.Register(http.StatusInternalServerError, client.UnknownError)
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := &dto.CreateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}

	res, err := h.clientUcase.Create(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, client.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) GetReservations(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &rdto.GetAllByClientReq{ClientID: p.ByName("id")}

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err != nil {
		httpjson.InternalError(w, nil)
		return
	} else if claims.UserID != req.ClientID {
		httpjson.Forbidden(w, nil)
		return
	}

	res, err := h.reservationUcase.GetAllByClient(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, client.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

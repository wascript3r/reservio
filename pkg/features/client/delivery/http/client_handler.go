package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errutil"
	"github.com/wascript3r/reservio/pkg/features/client"
	"github.com/wascript3r/reservio/pkg/features/client/dto"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	rdto "github.com/wascript3r/reservio/pkg/features/reservation/dto"
)

var parseErr = errutil.ParseErrFunc(client.InvalidInputError, client.UnknownError)

type HTTPHandler struct {
	clientUcase      client.Usecase
	reservationUcase reservation.Usecase
}

func NewHTTPHandler(r *httprouter.Router, cu client.Usecase, ru reservation.Usecase) {
	handler := &HTTPHandler{
		clientUcase:      cu,
		reservationUcase: ru,
	}

	r.POST("/client", handler.Create)
	r.GET("/client/:id/reservations", handler.GetReservations)
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
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) GetReservations(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &rdto.GetAllByClientReq{ClientID: p.ByName("id")}

	res, err := h.reservationUcase.GetAllByClient(r.Context(), req)
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

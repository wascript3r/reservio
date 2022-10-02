package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errutil"
	"github.com/wascript3r/reservio/pkg/features/reservation"
	"github.com/wascript3r/reservio/pkg/features/reservation/dto"
)

const InitRoute = "/company/:companyID/service/:serviceID"

var parseErr = errutil.ParseErrFunc(reservation.InvalidInputError, reservation.UnknownError)

type HTTPHandler struct {
	reservationUcase reservation.Usecase
}

func NewHTTPHandler(r *httprouter.Router, ru reservation.Usecase) {
	handler := &HTTPHandler{
		reservationUcase: ru,
	}

	r.POST(InitRoute+"/reservation", handler.Create)
	r.GET(InitRoute+"/reservation/:reservationID", handler.Get)
	// r.GET(InitRoute+"/reservations", handler.GetAll)
	// r.PATCH(InitRoute+"/reservation/:reservationID", handler.Update)
	// r.DELETE(InitRoute+"/reservation/:reservationID", handler.Delete)
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.CreateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")

	res, err := h.reservationUcase.Create(r.Context(), req)
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.GetReq{}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")
	req.ReservationID = p.ByName("reservationID")

	res, err := h.reservationUcase.Get(r.Context(), req, false)
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

// func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	req := &dto.GetAllReq{}
// 	req.CompanyID = p.ByName("companyID")
// 	req.ServiceID = p.ByName("serviceID")
//
// 	res, err := h.reservationUcase.GetAll(r.Context(), req, false)
// 	if err != nil {
// 		et, code := parseErr(err)
// 		errutil.ServeHTTP(w, et, code)
// 		return
// 	}
//
// 	httpjson.ServeJSON(w, res)
// }
//
// func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	req := &dto.UpdateReq{}
//
// 	err := json.NewDecoder(r.Body).Decode(req)
// 	if err != nil {
// 		httpjson.BadRequest(w, nil)
// 		return
// 	}
// 	req.CompanyID = p.ByName("companyID")
// 	req.ServiceID = p.ByName("serviceID")
// 	req.ReservationID = p.ByName("reservationID")
//
// 	err = h.reservationUcase.Update(r.Context(), req)
// 	if err != nil {
// 		et, code := parseErr(err)
// 		errutil.ServeHTTP(w, et, code)
// 		return
// 	}
//
// 	httpjson.ServeJSON(w, nil)
// }
//
// func (h *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	req := &dto.DeleteReq{}
// 	req.CompanyID = p.ByName("companyID")
// 	req.ServiceID = p.ByName("serviceID")
// 	req.ReservationID = p.ByName("reservationID")
//
// 	err := h.reservationUcase.Delete(r.Context(), req)
// 	if err != nil {
// 		et, code := parseErr(err)
// 		errutil.ServeHTTP(w, et, code)
// 		return
// 	}
//
// 	httpjson.ServeJSON(w, nil)
// }

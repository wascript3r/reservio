package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/features/service/dto"
)

const InitRoute = "/company/:companyID"

type HTTPHandler struct {
	mapper       *httpjson.CodeMapper
	serviceUcase service.Usecase
}

func NewHTTPHandler(r *httprouter.Router, cm *httpjson.CodeMapper, su service.Usecase) {
	handler := &HTTPHandler{
		mapper:       cm,
		serviceUcase: su,
	}
	handler.initErrs()

	r.POST(InitRoute+"/service", handler.Create)
	r.GET(InitRoute+"/service/:serviceID", handler.Get)
	r.GET(InitRoute+"/services", handler.GetAll)
	r.PATCH(InitRoute+"/service/:serviceID", handler.Update)
	r.DELETE(InitRoute+"/service/:serviceID", handler.Delete)
}

func (h *HTTPHandler) initErrs() {
	h.mapper.Register(http.StatusBadRequest, service.InvalidInputError, service.NothingToUpdateError)
	h.mapper.Register(http.StatusNotFound, service.NotFoundError)
	h.mapper.Register(http.StatusInternalServerError, service.UnknownError)
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.CreateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}
	req.CompanyID = p.ByName("companyID")

	res, err := h.serviceUcase.Create(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, service.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.GetReq{}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")

	res, err := h.serviceUcase.Get(r.Context(), req, false)
	if err != nil {
		code := errcode.UnwrapErr(err, service.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.GetAllReq{CompanyID: p.ByName("companyID")}

	res, err := h.serviceUcase.GetAll(r.Context(), req, false)
	if err != nil {
		code := errcode.UnwrapErr(err, service.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.UpdateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")

	err = h.serviceUcase.Update(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, service.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

func (h *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.DeleteReq{}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")

	err := h.serviceUcase.Delete(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, service.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

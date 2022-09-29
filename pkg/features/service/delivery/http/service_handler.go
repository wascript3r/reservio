package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errutil"
	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/features/service/dto"
)

const InitRoute = "/company/:companyID"

var parseErr = errutil.ParseErrFunc(service.InvalidInputError, service.UnknownError)

type HTTPHandler struct {
	serviceUcase service.Usecase
}

func NewHTTPHandler(r *httprouter.Router, su service.Usecase) {
	handler := &HTTPHandler{
		serviceUcase: su,
	}

	r.POST(InitRoute+"/service", handler.Create)
	r.GET(InitRoute+"/service/:serviceID", handler.Get)
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

	res, err := h.serviceUcase.Get(r.Context(), req, false)
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

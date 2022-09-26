package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errutil"
	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/company/dto"
)

var parseErr = errutil.ParseErrFunc(company.InvalidInputError, company.UnknownError)

type HTTPHandler struct {
	companyUcase company.Usecase
}

func NewHTTPHandler(r *httprouter.Router, cu company.Usecase) {
	handler := &HTTPHandler{
		companyUcase: cu,
	}

	r.POST("/company", handler.Create)
	r.GET("/company/:companyID", handler.Get)
	r.GET("/companies", handler.GetAll)
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := &dto.CreateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}

	res, err := h.companyUcase.Create(r.Context(), req)
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.GetReq{CompanyID: p.ByName("companyID")}
	res, err := h.companyUcase.Get(r.Context(), req)
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := h.companyUcase.GetAll(r.Context())
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, res)
}

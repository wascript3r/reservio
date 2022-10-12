package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
	"github.com/wascript3r/reservio/pkg/features/company"
	"github.com/wascript3r/reservio/pkg/features/company/dto"
	"github.com/wascript3r/reservio/pkg/features/token"
	mid "github.com/wascript3r/reservio/pkg/features/token/delivery/http"
)

const InitRoute = "/v1/companies"

type HTTPHandler struct {
	mapper       *httpjson.CodeMapper
	companyUcase company.Usecase
	tokenUcase   token.Usecase
}

func NewHTTPHandler(ctx context.Context, r *httprouter.Router, company mid.Company, admin mid.Admin, mp *httpjson.CodeMapper, cu company.Usecase, tu token.Usecase) {
	handler := &HTTPHandler{
		mapper:       mp,
		companyUcase: cu,
		tokenUcase:   tu,
	}
	handler.initErrs()

	r.POST(InitRoute, handler.Create)
	r.GET(InitRoute+"/:companyID", handler.Get)
	r.GET(InitRoute, handler.GetAll)
	r.PATCH(InitRoute+"/:companyID", company.Wrap(ctx, handler.Update))
	r.DELETE(InitRoute+"/:companyID", admin.Wrap(ctx, handler.Delete))
}

func (h *HTTPHandler) initErrs() {
	h.mapper.Register(http.StatusBadRequest, company.InvalidInputError, company.NothingToUpdateError)
	h.mapper.Register(http.StatusNotFound, company.NotFoundError)
	h.mapper.Register(http.StatusInternalServerError, company.UnknownError)
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
		code := errcode.UnwrapErr(err, company.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.GetReq{CompanyID: p.ByName("companyID")}
	res, err := h.companyUcase.Get(r.Context(), req, false)
	if err != nil {
		code := errcode.UnwrapErr(err, company.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := h.companyUcase.GetAll(r.Context(), false)
	if err != nil {
		code := errcode.UnwrapErr(err, company.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

func (h *HTTPHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.UpdateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}
	req.CompanyID = p.ByName("companyID")

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err != nil {
		httpjson.InternalError(w, nil)
		return
	} else if claims.UserID != req.CompanyID {
		httpjson.Forbidden(w, nil)
		return
	}

	err = h.companyUcase.Update(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, company.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

func (h *HTTPHandler) Delete(_ context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.DeleteReq{CompanyID: p.ByName("companyID")}
	err := h.companyUcase.Delete(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, company.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

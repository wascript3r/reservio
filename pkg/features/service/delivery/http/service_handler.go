package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
	"github.com/wascript3r/reservio/pkg/features/service"
	"github.com/wascript3r/reservio/pkg/features/service/dto"
	"github.com/wascript3r/reservio/pkg/features/token"
	mid "github.com/wascript3r/reservio/pkg/features/token/delivery/http"
)

const InitRoute = "/api/v1/companies/:companyID/services"

type HTTPHandler struct {
	mapper       *httpjson.CodeMapper
	serviceUcase service.Usecase
	tokenUcase   token.Usecase
}

func NewHTTPHandler(ctx context.Context, r *httprouter.Router, company mid.Company, cm *httpjson.CodeMapper, su service.Usecase, tu token.Usecase) {
	handler := &HTTPHandler{
		mapper:       cm,
		serviceUcase: su,
		tokenUcase:   tu,
	}
	handler.initErrs()

	r.POST(InitRoute, company.Wrap(ctx, handler.Create))
	r.GET(InitRoute+"/:serviceID", handler.Get)
	r.GET(InitRoute, handler.GetAll)
	r.PATCH(InitRoute+"/:serviceID", company.Wrap(ctx, handler.Update))
	r.DELETE(InitRoute+"/:serviceID", company.Wrap(ctx, handler.Delete))
}

func (h *HTTPHandler) initErrs() {
	h.mapper.Register(http.StatusBadRequest, service.InvalidInputError, service.NothingToUpdateError)
	h.mapper.Register(http.StatusNotFound, service.NotFoundError)
	h.mapper.Register(http.StatusInternalServerError, service.UnknownError)
}

func (h *HTTPHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.CreateReq{}

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

func (h *HTTPHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.UpdateReq{}

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
	} else if claims.UserID != req.CompanyID {
		httpjson.Forbidden(w, nil)
		return
	}

	err = h.serviceUcase.Update(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, service.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

func (h *HTTPHandler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req := &dto.DeleteReq{}
	req.CompanyID = p.ByName("companyID")
	req.ServiceID = p.ByName("serviceID")

	claims, err := h.tokenUcase.LoadCtx(ctx)
	if err != nil {
		httpjson.InternalError(w, nil)
		return
	} else if claims.UserID != req.CompanyID {
		httpjson.Forbidden(w, nil)
		return
	}

	err = h.serviceUcase.Delete(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, service.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.Status(w, http.StatusNoContent)
}

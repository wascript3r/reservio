package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
	"github.com/wascript3r/reservio/pkg/features/token"
	"github.com/wascript3r/reservio/pkg/features/token/dto"
)

const InitRoute = "/api/v1/tokens"

type HTTPHandler struct {
	mapper     *httpjson.CodeMapper
	tokenUcase token.Usecase
}

func NewHTTPHandler(r *httprouter.Router, mp *httpjson.CodeMapper, tu token.Usecase) {
	handler := &HTTPHandler{
		mapper:     mp,
		tokenUcase: tu,
	}
	handler.initErrs()

	r.POST(InitRoute, handler.RenewAccess)
}

func (h *HTTPHandler) initErrs() {
	h.mapper.Register(http.StatusBadRequest, token.InvalidInputError, token.FaultyTokenError)
	h.mapper.Register(http.StatusUnauthorized, token.InvalidOrExpiredTokenError)
	h.mapper.Register(http.StatusInternalServerError, token.UnknownError)
}

func (h *HTTPHandler) RenewAccess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := &dto.RenewAccessReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}

	res, err := h.tokenUcase.RenewAccess(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, token.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

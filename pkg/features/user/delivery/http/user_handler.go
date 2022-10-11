package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
	"github.com/wascript3r/reservio/pkg/features/user"
	"github.com/wascript3r/reservio/pkg/features/user/dto"
)

const InitRoute = "/api/v1/users"

type HTTPHandler struct {
	mapper    *httpjson.CodeMapper
	userUcase user.Usecase
}

func NewHTTPHandler(r *httprouter.Router, cm *httpjson.CodeMapper, uu user.Usecase) {
	handler := &HTTPHandler{
		mapper:    cm,
		userUcase: uu,
	}
	handler.initErrs()

	r.POST(InitRoute+"/authenticate", handler.Authenticate)
}

func (h *HTTPHandler) initErrs() {
	h.mapper.Register(http.StatusBadRequest, user.InvalidInputError)
	h.mapper.Register(http.StatusUnprocessableEntity, user.EmailAlreadyExistsError)
	h.mapper.Register(http.StatusUnauthorized, user.InvalidCredentialsError)
	h.mapper.Register(http.StatusInternalServerError, user.UnknownError)
}

func (h *HTTPHandler) Authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := &dto.AuthenticateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}

	res, err := h.userUcase.Authenticate(r.Context(), req)
	if err != nil {
		code := errcode.UnwrapErr(err, user.UnknownError)
		h.mapper.ServeErr(w, code, nil)
		return
	}

	httpjson.ServeJSON(w, res)
}

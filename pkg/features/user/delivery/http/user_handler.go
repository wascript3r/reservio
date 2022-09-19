package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errutil"
	"github.com/wascript3r/reservio/pkg/features/user"
	udto "github.com/wascript3r/reservio/pkg/features/user/dto"
)

var parseErr = errutil.ParseErrFunc(user.InvalidInputError, user.UnknownError)

type HTTPHandler struct {
	userUcase user.Usecase
}

func NewHTTPHandler(r *httprouter.Router, uu user.Usecase) {
	handler := &HTTPHandler{
		userUcase: uu,
	}

	r.POST("/user/create", handler.Create)
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := &udto.CreateReq{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		httpjson.BadRequest(w, nil)
		return
	}

	err = h.userUcase.Create(r.Context(), req)
	if err != nil {
		et, code := parseErr(err)
		errutil.ServeHTTP(w, et, code)
		return
	}

	httpjson.ServeJSON(w, nil)
}

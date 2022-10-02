package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errutil"
	"github.com/wascript3r/reservio/pkg/features/client"
	"github.com/wascript3r/reservio/pkg/features/client/dto"
)

var parseErr = errutil.ParseErrFunc(client.InvalidInputError, client.UnknownError)

type HTTPHandler struct {
	clientUcase client.Usecase
}

func NewHTTPHandler(r *httprouter.Router, cu client.Usecase) {
	handler := &HTTPHandler{
		clientUcase: cu,
	}

	r.POST("/client", handler.Create)
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

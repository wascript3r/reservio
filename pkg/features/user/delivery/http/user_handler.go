package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wascript3r/reservio/pkg/features/user"
)

// var parseErr = errutil.ParseErrFunc(user.InvalidInputError, user.UnknownError)

type HTTPHandler struct {
	userUcase user.Usecase
}

func NewHTTPHandler(r *httprouter.Router, uu user.Usecase) {
	// handler := &HTTPHandler{
	// 	userUcase: uu,
	// }

}

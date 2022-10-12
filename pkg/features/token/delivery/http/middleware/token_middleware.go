package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/wascript3r/httputil"
	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/features/token"
	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

var (
	ErrCannotExtractToken = errors.New("cannot extract token")
)

type HTTPMiddleware struct {
	tokenUcase token.Usecase
}

func NewHTTPMiddleware(tu token.Usecase) *HTTPMiddleware {
	return &HTTPMiddleware{
		tokenUcase: tu,
	}
}

func (h *HTTPMiddleware) ExtractToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", ErrCannotExtractToken
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrCannotExtractToken
	}

	return parts[1], nil
}

func (h *HTTPMiddleware) Authenticated(next httputil.HandleCtx) httputil.HandleCtx {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		t, err := h.ExtractToken(r)
		if err != nil {
			httpjson.Unauthorized(w, nil)
			return
		}

		claims, err := h.tokenUcase.ParseAccess(t)
		if err != nil {
			httpjson.UnauthorizedCustom(w, token.InvalidOrExpiredTokenError, nil)
			return
		}
		ctx = h.tokenUcase.StoreCtx(ctx, claims)

		next(ctx, w, r, p)
	}
}

func (h *HTTPMiddleware) ParseUser(next httputil.HandleCtx) httputil.HandleCtx {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		t, err := h.ExtractToken(r)
		if err != nil {
			next(ctx, w, r, p)
			return
		}

		claims, err := h.tokenUcase.ParseAccess(t)
		if err != nil {
			httpjson.UnauthorizedCustom(w, token.InvalidOrExpiredTokenError, nil)
			return
		}
		ctx = h.tokenUcase.StoreCtx(ctx, claims)

		next(ctx, w, r, p)
	}
}

func roleExists(roles []umodels.Role, role umodels.Role) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func (h *HTTPMiddleware) IsOneOf(role ...umodels.Role) func(next httputil.HandleCtx) httputil.HandleCtx {
	return func(next httputil.HandleCtx) httputil.HandleCtx {
		return h.Authenticated(
			func(ctx context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				claims, err := h.tokenUcase.LoadCtx(ctx)
				if err != nil {
					httpjson.Unauthorized(w, nil)
					return
				}

				if !roleExists(role, claims.Role) {
					httpjson.Forbidden(w, nil)
					return
				}

				next(ctx, w, r, p)
			},
		)
	}
}

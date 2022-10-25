package middleware

import (
	"net/http"
)

type HTTPMiddleware struct {
	allowOrigin string
}

func NewHTTPMiddleware(allowOrigin string) *HTTPMiddleware {
	return &HTTPMiddleware{allowOrigin}
}

func (h *HTTPMiddleware) EnableCors(hnd http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", h.allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		hnd.ServeHTTP(w, r)
	})
}

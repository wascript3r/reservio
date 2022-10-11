package http

import (
	"github.com/wascript3r/httputil"
	"github.com/wascript3r/httputil/middleware"
	"github.com/wascript3r/reservio/pkg/features/user/models"
)

type Middleware interface {
	Authenticated(next httputil.HandleCtx) httputil.HandleCtx
	HasRole(role models.Role) func(next httputil.HandleCtx) httputil.HandleCtx
}

type (
	Auth    struct{ *middleware.StackCtx }
	Admin   struct{ *middleware.StackCtx }
	Company struct{ *middleware.StackCtx }
	Client  struct{ *middleware.StackCtx }
)

package main

import (
	"errors"
	"net/http"

	httpjson "github.com/wascript3r/httputil/json"
	"github.com/wascript3r/reservio/pkg/errcode"
)

var (
	MethodNotAllowedError = errcode.New(
		"method_not_allowed",
		errors.New("method not allowed"),
	)
	MethodNotAllowedHnd = http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			httpjson.BadRequestCustom(w, MethodNotAllowedError, nil)
		},
	)

	NotFoundHnd = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			httpjson.NotFound(w, nil)
		},
	)
)

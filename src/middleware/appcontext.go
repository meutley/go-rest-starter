package middleware

import (
	"context"
	"net/http"

	"../common"
)

func AppContext(h common.RouteHandler, ctx context.Context) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return h(w, r.WithContext(ctx))
	}
}

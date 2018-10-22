package router

import (
	"context"
	"net/http"

	"../database"
)

func appContext(h RouteHandler, ctx context.Context) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return h(w, r.WithContext(ctx))
	}
}

// GetDatabaseFromContext gets a DatabaseContract interface from the request Context.
func GetDatabaseFromContext(r *http.Request) database.DatabaseContract {
	return r.Context().Value("DatabaseContract").(database.DatabaseContract)
}

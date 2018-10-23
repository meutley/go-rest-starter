package common

import (
	"net/http"

	"../database"
)

// GetDatabaseFromContext gets a DatabaseContract interface from the request Context.
func GetDatabaseFromContext(r *http.Request) database.DatabaseContract {
	return r.Context().Value("DatabaseContract").(database.DatabaseContract)
}

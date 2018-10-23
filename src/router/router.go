package router

import (
	"context"
	"net/http"
	"path"

	"../common"
	"../database"
	"../middleware"
	"github.com/gorilla/mux"
)

func applyMiddleware(h common.RouteHandler, db database.DatabaseContract) func(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "DatabaseContract", db)
	return middleware.ErrorWrapper(middleware.AppContext(h, ctx))
}

// Route defines an API route via various configurable parameters.
type Route struct {
	Name              string
	Method            string
	RootPattern       string
	PatternParams     string
	HandlerFunc       common.RouteHandler
	ChildRoutes       []Route
	WithParentPattern string
}

// Routes represents a slice of Route objects.
type Routes []Route

// BuildRouter builds a new instance of mux.Router using the API route handlers.
func BuildRouter(routes Routes, db database.DatabaseContract) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()

	for _, parentRoute := range routes {
		// Parent route handler
		fullPattern := path.Join(parentRoute.RootPattern, parentRoute.PatternParams)
		sub.
			HandleFunc(fullPattern, applyMiddleware(parentRoute.HandlerFunc, db)).
			Name(parentRoute.Name).
			Methods(parentRoute.Method)

		// Build child routes
		for _, child := range parentRoute.ChildRoutes {
			childPattern := path.Join(parentRoute.RootPattern,
				child.WithParentPattern,
				child.RootPattern,
				child.PatternParams)
			childName := path.Join(parentRoute.Name, child.Name)

			sub.
				HandleFunc(childPattern, applyMiddleware(child.HandlerFunc, db)).
				Name(childName).
				Methods(child.Method)
		}
	}

	router.Use(middleware.Json)
	return router
}

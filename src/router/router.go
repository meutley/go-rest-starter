package router

import (
	"path"

	"../middleware"
	"github.com/gorilla/mux"
)

// BuildRouter builds a new instance of mux.Router using the API route handlers.
func BuildRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()

	for _, route := range routes {
		// Parent route handler
		fullPattern := path.Join(route.RootPattern, route.PatternParams)
		sub.
			HandleFunc(fullPattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)

		// Build child routes
		for _, child := range route.ChildRoutes {
			childPattern := path.Join(route.RootPattern, child.WithParentPattern, child.RootPattern)
			childName := path.Join(route.Name, child.Name)

			sub.
				HandleFunc(childPattern, child.HandlerFunc).
				Name(childName).
				Methods(child.Method)
		}
	}

	router.Use(middleware.Json)
	return router
}

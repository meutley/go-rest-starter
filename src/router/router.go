package router

import (
	"path"

	"../middleware"
	"github.com/gorilla/mux"
)

// BuildRouter builds a new instance of mux.Router using the API route handlers.
func BuildRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()

	for _, parentRoute := range routes {
		// Parent route handler
		fullPattern := path.Join(parentRoute.RootPattern, parentRoute.PatternParams)
		sub.
			HandleFunc(fullPattern, errorWrapper(parentRoute.HandlerFunc)).
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
				HandleFunc(childPattern, errorWrapper(child.HandlerFunc)).
				Name(childName).
				Methods(child.Method)
		}
	}

	router.Use(middleware.Json)
	return router
}

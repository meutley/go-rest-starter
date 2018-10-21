package router

import (
	"net/http"

	"../api"
)

// Route defines an API route via various configurable parameters.
type Route struct {
	Name              string
	Method            string
	RootPattern       string
	PatternParams     string
	HandlerFunc       http.HandlerFunc
	ChildRoutes       []Route
	WithParentPattern string
}

// Routes represents a slice of Route objects.
type Routes []Route

var AppRoutes = Routes{
	// GET /api/v1/root
	Route{
		"GetIndex",
		"GET",
		"/root",
		"",
		api.GetRoot,
		Routes{
			// GET /api/v1/root/{root_id}/test
			Route{
				"GetTest",
				"GET",
				"/test",
				"",
				api.GetRootTest,
				Routes{},
				"{root_id}",
			},
		},
		"",
	},
}

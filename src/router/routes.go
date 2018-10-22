package router

import (
	"net/http"
)

// RouteHandler is a custom route handler function that can return an error.
type RouteHandler func(w http.ResponseWriter, r *http.Request) error

// Route defines an API route via various configurable parameters.
type Route struct {
	Name              string
	Method            string
	RootPattern       string
	PatternParams     string
	HandlerFunc       RouteHandler
	ChildRoutes       []Route
	WithParentPattern string
}

// Routes represents a slice of Route objects.
type Routes []Route

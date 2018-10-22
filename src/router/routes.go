package router

import (
	"net/http"
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

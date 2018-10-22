package root

import (
	"errors"
	"net/http"

	"../../common"
	"../../router"
)

func get400Error(w http.ResponseWriter, r *http.Request) error {
	return &common.StatusError{
		Status: 400,
		Err:    errors.New("api/root/bad/400 generates a 400 error response"),
	}
}

func get404Error(w http.ResponseWriter, r *http.Request) error {
	return &common.StatusError{
		Status: 404,
		Err:    errors.New("api/root/bad/400 generates a 404 error response"),
	}
}

func get500Error(w http.ResponseWriter, r *http.Request) error {
	return &common.StatusError{
		Status: 500,
		Err:    errors.New("api/root/bad/400 generates a 500 error response"),
	}
}

var badRoutes = router.Routes{
	// GET /api/root/bad/400
	router.Route{
		Name:              "Get400Error",
		Method:            "GET",
		RootPattern:       "/bad",
		PatternParams:     "400",
		HandlerFunc:       get400Error,
		ChildRoutes:       router.Routes{},
		WithParentPattern: "",
	},
	// GET /api/root/bad/404
	router.Route{
		Name:              "Get404Error",
		Method:            "GET",
		RootPattern:       "/bad",
		PatternParams:     "404",
		HandlerFunc:       get404Error,
		ChildRoutes:       router.Routes{},
		WithParentPattern: "",
	},
	// GET /api/root/bad/500
	router.Route{
		Name:              "Get500Error",
		Method:            "GET",
		RootPattern:       "/bad",
		PatternParams:     "500",
		HandlerFunc:       get500Error,
		ChildRoutes:       router.Routes{},
		WithParentPattern: "",
	},
}

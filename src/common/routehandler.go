package common

import "net/http"

type RouteHandler func(w http.ResponseWriter, r *http.Request) error

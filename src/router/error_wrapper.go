package router

import (
	"net/http"

	"../common"
)

func errorWrapper(h RouteHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			switch e := err.(type) {
			case *common.StatusError:
				switch e.Status {
				case 400:
					http.Error(w, e.Err.Error(), 400)
					return
				case 404:
					http.NotFound(w, r)
					return
				default:
					http.Error(w, http.StatusText(500), 500)
					return
				}
			default:
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}
	}
}

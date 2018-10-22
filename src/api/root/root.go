package root

import (
	"encoding/json"
	"net/http"

	"../../router"
)

func getRoot(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
	return nil
}

var Routes = router.Routes{
	// GET /api/root
	router.Route{
		Name:              "GetRoot",
		Method:            "GET",
		RootPattern:       "/root",
		PatternParams:     "",
		HandlerFunc:       getRoot,
		ChildRoutes:       append(nestedRoutes, badRoutes...),
		WithParentPattern: "",
	},
}

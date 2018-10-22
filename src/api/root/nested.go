package root

import (
	"encoding/json"
	"net/http"

	"../../router"
)

func getNested(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("root / nested")
	return nil
}

var nestedRoutes = router.Routes{
	// GET /api/root/{root_id}/nested
	router.Route{
		Name:              "GetNested",
		Method:            "GET",
		RootPattern:       "/nested",
		PatternParams:     "",
		HandlerFunc:       getNested,
		ChildRoutes:       router.Routes{},
		WithParentPattern: "{root_id}",
	},
}

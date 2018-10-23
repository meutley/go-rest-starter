package root

import (
	"encoding/json"
	"net/http"

	"../../common"
	"../../router"
)

func getRoot(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
	return nil
}

func getUser1(w http.ResponseWriter, r *http.Request) error {
	db := common.GetDatabaseFromContext(r)
	u, err := db.FindUser(1)
	if err != nil {
		return err
	} else if u.Id == 0 {
		return &common.StatusError{Status: 404, Err: nil}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
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
	// GET /api/root/getUser1
	router.Route{
		Name:              "GetRoot",
		Method:            "GET",
		RootPattern:       "/root/getUser1",
		PatternParams:     "",
		HandlerFunc:       getUser1,
		ChildRoutes:       router.Routes{},
		WithParentPattern: "",
	},
}

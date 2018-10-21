package api

import (
	"encoding/json"
	"net/http"
)

func GetRootTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("root / test")
}

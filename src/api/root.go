package api

import (
	"encoding/json"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

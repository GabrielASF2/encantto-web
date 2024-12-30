package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ParseJSONRequest(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(target)
}

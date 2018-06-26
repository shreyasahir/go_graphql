package app

import (
	"encoding/json"
	"net/http"
)

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

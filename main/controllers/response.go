package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
}

func (r Response) withError(w http.ResponseWriter, code int, message string) {
	r.withJSON(w, code, map[string]string{"error": message})
}

func (r Response) withJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

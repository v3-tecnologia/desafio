package web

import (
	"encoding/json"
	"net/http"
)

// Error represents an error response
type Error struct {
	Message string `json:"message"`
}

// HandleError handles the error response
func HandleError(w http.ResponseWriter, err error, validationErrors map[string]bool) {
	status := http.StatusInternalServerError
	if validationErrors[err.Error()] {
		status = http.StatusBadRequest
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Error{Message: err.Error()})
}

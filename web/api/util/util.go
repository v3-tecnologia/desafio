package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Warn(err)
	}
}

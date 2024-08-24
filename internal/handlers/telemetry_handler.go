package handlers

import (
	"github.com/ThalesMonteir0/desafio/internal/ports"
	"net/http"
)

type telemetryHandler struct {
	service ports.TelemetryService
}

type TelemetryHandler interface {
	CreateGyroscope(writer http.ResponseWriter, request *http.Request)
	CreateGPS(writer http.ResponseWriter, request *http.Request)
	CreatePhoto(writer http.ResponseWriter, request *http.Request)
}

func NewTelemetryHandler(service ports.TelemetryService) TelemetryHandler {
	return &telemetryHandler{
		service: service,
	}
}

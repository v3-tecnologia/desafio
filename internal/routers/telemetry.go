package routers

import (
	"github.com/ThalesMonteir0/desafio/internal/handlers"
	"net/http"
)

func TelemetryRoutes(server *http.ServeMux, telemetryHandler handlers.TelemetryHandler) {

	server.HandleFunc("POST /telemetry/gyroscope", telemetryHandler.CreateGyroscope)
	server.HandleFunc("POST /telemetry/gps", telemetryHandler.CreateGPS)
	server.HandleFunc("POST /telemetry/photo", telemetryHandler.CreatePhoto)

}

package api

import (
	"desafio-backend/web/api/v1/telemetry"
	"fmt"
	"github.com/gorilla/mux"
)

func (api *API) newV1Api(router *mux.Router) {
	s := router.PathPrefix("/v1").Subrouter()
	{
		pathTelemetry := "/telemetry"
		s.HandleFunc(fmt.Sprintf("%s/gyroscope", pathTelemetry), telemetry.Gyroscope(api.gyroscopeMain)).Methods("POST")
		s.HandleFunc(fmt.Sprintf("%s/gps", pathTelemetry), telemetry.Gps()).Methods("POST")
		s.HandleFunc(fmt.Sprintf("%s/photo", pathTelemetry), telemetry.Photo()).Methods("POST")
	}
}

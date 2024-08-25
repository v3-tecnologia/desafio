package handlers

import (
	"encoding/json"
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/handlers/models/request"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
	"io"
	"net/http"
)

func (t telemetryHandler) CreateGPS(w http.ResponseWriter, r *http.Request) {
	var gpsRequest request.GpsRequest
	var restErr *err_rest.ErrRest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		restErr = err_rest.NewInternalServerError("unable read body")
		w.WriteHeader(restErr.Code)
		return
	}

	if err = json.Unmarshal(body, &gpsRequest); err != nil {
		restErr = err_rest.NewUnprocessableEntityError("unable to process the request because it contains invalid data")
		w.WriteHeader(restErr.Code)
		return
	}

	if restErr = validateBodyRequired(gpsRequest.FieldsRequiredTelemetry); restErr != nil {
		w.WriteHeader(restErr.Code)
		return
	}

	gpsDomain := domain.GpsDomain{
		Latitude:   gpsRequest.Latitude,
		Longitude:  gpsRequest.Longitude,
		MacAddress: gpsRequest.MacAddress,
	}

	if restErr = t.service.CreateGpsService(gpsDomain); restErr != nil {
		w.WriteHeader(restErr.Code)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

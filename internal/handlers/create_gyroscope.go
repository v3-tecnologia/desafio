package handlers

import (
	"encoding/json"
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/handlers/models/request"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
	"io"
	"net/http"
)

func (t telemetryHandler) CreateGyroscope(w http.ResponseWriter, r *http.Request) {
	var gyroscopeRequest request.GyroscopeRequest
	var errRest *err_rest.ErrRest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errRest = err_rest.NewBadRequestErr("unable read body")
		w.WriteHeader(errRest.Code)
		return
	}

	if err = json.Unmarshal(body, &gyroscopeRequest); err != nil {
		errRest = err_rest.NewUnprocessableEntityError("unable to process the request because it contains invalid data")
		w.WriteHeader(errRest.Code)
		return
	}

	if errRest = validateGyroscopeBody(gyroscopeRequest); err != nil {
		w.WriteHeader(errRest.Code)
		return
	}

	gyroscopeDomain := domain.GyroscopeDomain{
		X:          gyroscopeRequest.X,
		Y:          gyroscopeRequest.Y,
		Z:          gyroscopeRequest.Z,
		MacAddress: gyroscopeRequest.MacAddress,
	}

	if errRest = t.service.CreateGyroscopeService(gyroscopeDomain); err != nil {
		w.WriteHeader(errRest.Code)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func validateGyroscopeBody(gyroscopeRequest request.GyroscopeRequest) *err_rest.ErrRest {
	if gyroscopeRequest.MacAddress == "" {
		return err_rest.NewBadRequestErr("field mac_address is required")
	}

	if gyroscopeRequest.CollectionDate.IsZero() {
		return err_rest.NewBadRequestErr("field collection_date is required")
	}

	return nil
}

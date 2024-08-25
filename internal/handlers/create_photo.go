package handlers

import (
	"encoding/json"
	"github.com/ThalesMonteir0/desafio/internal/domain"
	"github.com/ThalesMonteir0/desafio/internal/handlers/models/request"
	"github.com/ThalesMonteir0/desafio/pkg/err_rest"
	"io"
	"net/http"
)

func (t telemetryHandler) CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var photoRequest request.PhotoRequest
	var restErr *err_rest.ErrRest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		restErr = err_rest.NewInternalServerError("unable read body")
		w.WriteHeader(restErr.Code)
		return
	}

	if err = json.Unmarshal(body, &photoRequest); err != nil {
		restErr = err_rest.NewUnprocessableEntityError("unable to process the request because it contains invalid data")
		w.WriteHeader(restErr.Code)
		return
	}

	if restErr = validatePhotoBody(photoRequest); restErr != nil {
		w.WriteHeader(restErr.Code)
		return
	}

	photoDomain := domain.PhotoDomain{
		Url:            photoRequest.Url,
		MacAddress:     photoRequest.MacAddress,
		CollectionDate: photoRequest.CollectionDate,
	}

	if restErr = t.service.CreatePhotoService(photoDomain); restErr != nil {
		w.WriteHeader(restErr.Code)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func validatePhotoBody(photoRequest request.PhotoRequest) *err_rest.ErrRest {
	if err := validateBodyRequired(photoRequest.FieldsRequiredTelemetry); err != nil {
		return err
	}

	if photoRequest.Url == "" {
		return err_rest.NewBadRequestErr("photo url is required")
	}

	return nil
}

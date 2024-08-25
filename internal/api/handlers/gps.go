package handlers

import (
	"errors"
	"net/http"
	"time"
	"v3/pkg/httpcore"
	"v3/pkg/models"
)

func (tc *ApiController) CreateGPS(w http.ResponseWriter, r *http.Request) (any, int) {
	newGPS, err := httpcore.DecodeBody[models.GPS](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	if newGPS.DeviceData == nil {
		return httpcore.ErrBadRequest.With(errors.New("device cannot be nil")), http.StatusBadRequest
	}

	g, err := models.NewGPS(
		&models.DeviceData{
			MAC:       newGPS.DeviceData.MAC,
			Timestamp: time.Now(),
		},
		newGPS.Latitude,
		newGPS.Longitude,
	)

	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	if tc.db == nil {
		tc.db = make([]DataModel, 0)
	}

	tc.db = append(tc.db, g)

	return g, http.StatusCreated
}

package handlers

import (
	"errors"
	"net/http"
	"v3/pkg/httpcore"
	"v3/pkg/models"
)

func (tc *ApiController) CreateGyroscope(w http.ResponseWriter, r *http.Request) (any, int) {
	newGyro, err := httpcore.DecodeBody[models.Gyroscope](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	if newGyro.DeviceData == nil {
		return httpcore.ErrBadRequest.With(errors.New("device data cannot be nil")), http.StatusBadRequest
	}
	if newGyro.X == nil || newGyro.Y == nil || newGyro.Z == nil {
		return httpcore.ErrBadRequest.With(errors.New("gyroscope data cannot have nil values")), http.StatusBadRequest
	}

	g, err := models.NewGyroscope(
		newGyro.DeviceData,
		newGyro.X,
		newGyro.Y,
		newGyro.Z,
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

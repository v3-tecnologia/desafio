package handlers

import (
	"errors"
	"net/http"
	"time"
	"v3/pkg/httpcore"
	"v3/pkg/models"
)

// CreatePhoto godoc
// @Summary Create a new photo
// @Description Create a new photo with device data and photo values
// @Accept  json
// @Produce  json
// @Param   photo  body   models.PhotoData  true  "Photo data"
// @Success 201 {object} models.PhotoData
// @Failure 400 {object} httpcore.ApiError
// @Router /telemetry/photo [post]
func (tc *ApiController) CreatePhoto(w http.ResponseWriter, r *http.Request) (any, int) {
	newPData, err := httpcore.DecodeBody[models.PhotoData](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	if newPData.DeviceData == nil {
		return httpcore.ErrBadRequest.With(errors.New("device cannot be nil")), http.StatusBadRequest
	}

	p, err := models.NewPhotoData(
		&models.DeviceData{
			MAC:       newPData.MAC,
			Timestamp: time.Now(),
		},
		newPData.Path,
	)

	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	tc.db = append(tc.db, p)

	return p, http.StatusCreated
}

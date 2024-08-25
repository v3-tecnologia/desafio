package handlers

import (
	"net/http"
	"time"
	"v3/pkg/httpcore"
	"v3/pkg/models"
	"v3/pkg/utils"
)

func (tc *ApiController) CreatePhotoData(w http.ResponseWriter, r *http.Request) (any, int) {
	newPData, err := httpcore.DecodeBody[models.PhotoData](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	p, err := models.NewPhotoData(
		&models.DeviceData{
			MAC:       utils.GenerateMac(),
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

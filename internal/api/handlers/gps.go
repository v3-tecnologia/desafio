package handlers

import (
	"net/http"
	"time"
	"v3/pkg/httpcore"
	"v3/pkg/models"
	"v3/pkg/util"
)

func (tc *ApiController) CreateGPS(w http.ResponseWriter, r *http.Request) (any, int) {
	newGPS, err := httpcore.DecodeBody[models.GPS](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	g, err := models.NewGPS(
		&models.DeviceData{
			MAC:       util.GenerateMac(),
			Timestamp: time.Now(),
		},
		newGPS.Latitude,
		newGPS.Longitude,
	)

	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	tc.db = append(tc.db, g)

	return g, http.StatusCreated
}

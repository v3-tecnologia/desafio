package handlers

import (
	"net/http"
	"time"
	"v3/pkg/httpcore"
	"v3/pkg/models"
	"v3/pkg/util"
)

func (tc *ApiController) CreateGyroscope(w http.ResponseWriter, r *http.Request) (any, int) {
	newGyro, err := httpcore.DecodeBody[models.Gyroscope](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	g, err := models.NewGyroscope(
		&models.DeviceData{
			MAC:       util.GenerateMac(),
			Timestamp: time.Now(),
		},
		newGyro.X,
		newGyro.Y,
		newGyro.Z,
	)

	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	tc.db = append(tc.db, g)

	return g, http.StatusCreated
}

package presenter

import (
	"time"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func DomainGPSToHTTP(gps entity.GPS) HTTPGPS {
	return HTTPGPS{
		ID:        gps.ID,
		Latitude:  gps.Latitude,
		Longitude: gps.Longitude,
		CreatedAt: gps.CreatedAt,
	}

}

type HTTPGPS struct {
	ID        string    `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
}

package mapper

import (
	"time"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func DomainGPSToPostgres(gps entity.GPS) PostgresGPS {
	return PostgresGPS{
		ID:        gps.ID,
		Latitude:  gps.Latitude,
		Longitude: gps.Longitude,
		CreatedAt: gps.CreatedAt,
	}
}

func PostgresGPSToDomain(gps PostgresGPS) entity.GPS {
	return entity.GPS{
		ID:        gps.ID,
		Latitude:  gps.Latitude,
		Longitude: gps.Longitude,
		CreatedAt: gps.CreatedAt,
	}
}

type PostgresGPS struct {
	ID        string    `json:"id" db:"id"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

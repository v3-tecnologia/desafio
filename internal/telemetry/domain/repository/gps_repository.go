package repository

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

type GPSRepository interface {
	Store(gps entity.GPS) error
}

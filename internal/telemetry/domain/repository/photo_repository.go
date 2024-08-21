package repository

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

type PhotoRepository interface {
	Store(photo entity.Photo) error
}

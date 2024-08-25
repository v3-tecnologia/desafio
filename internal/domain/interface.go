package domain

import "github.com/HaroldoFV/desafio/internal/domain/entity"

type GyroscopeRepositoryInterface interface {
	Create(gyroscope *entity.Gyroscope) error
}

type GPSRepositoryInterface interface {
	Create(gps *entity.GPS) error
}

type PhotoRepositoryInterface interface {
	Create(photo *entity.Photo) error
}

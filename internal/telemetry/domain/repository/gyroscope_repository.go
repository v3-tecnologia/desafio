package repository

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

type GyroscopeRepository interface {
	Store(gyroscope entity.Gyroscope) error
}

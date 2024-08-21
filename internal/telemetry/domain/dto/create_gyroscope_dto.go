package dto

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

type CreateGyroscopeInputDTO struct {
	XPosition float64
	YPosition float64
	ZPosition float64
}

type CreateGyroscopeOutputDTO struct {
	Gyroscope entity.Gyroscope
}

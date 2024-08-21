package dto

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

type CreateGPSInputDTO struct {
	Latitude  float64
	Longitude float64
}

type CreateGPSOutputDTO struct {
	GPS entity.GPS
}

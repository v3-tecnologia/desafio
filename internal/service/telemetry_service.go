package service

import (
	"github.com/ThalesMonteir0/desafio/internal/ports"
)

type telemetryService struct {
	repository ports.TelemetryRepository
}

func NewTelemetryService(repository ports.TelemetryRepository) ports.TelemetryService {
	return &telemetryService{
		repository: repository,
	}
}

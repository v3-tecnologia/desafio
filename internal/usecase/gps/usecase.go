package gps

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/gps"
	"github.com/kevenmiano/v3/internal/usecase/gps/create"
)

type UseCase interface {
	Execute(d *domain.GPS) (*domain.GPS, error)
}

func NewCreateGpsUseCase(gpsRepository gps.Repository) UseCase {
	return &create.GpsUseCase{
		GpsRepository: gpsRepository,
	}
}

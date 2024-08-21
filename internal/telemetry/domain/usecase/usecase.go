package usecase

import (
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/domain/repository"
)

type TelemetryUseCase interface {
	CreatePhoto(input dto.CreatePhotoInputDTO) (dto.CreatePhotoOutputDTO, error)
	CreateGPS(input dto.CreateGPSInputDTO) (dto.CreateGPSOutputDTO, error)
	CreateGyroscope(input dto.CreateGyroscopeInputDTO) (dto.CreateGyroscopeOutputDTO, error)
}

func NewTelemetryUseCaseRegistry(
	gpsRepo repository.GPSRepository,
	gyroscopeRepo repository.GyroscopeRepository,
	photoRepo repository.PhotoRepository,
) TelemetryUseCaseRegistry {
	return TelemetryUseCaseRegistry{
		gpsRepo:       gpsRepo,
		gyroscopeRepo: gyroscopeRepo,
		photoRepo:     photoRepo,
	}
}

type TelemetryUseCaseRegistry struct {
	gpsRepo       repository.GPSRepository
	gyroscopeRepo repository.GyroscopeRepository
	photoRepo     repository.PhotoRepository
}

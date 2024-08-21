package usecase

import (
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/domain/port"
	"github.com/charmingruby/g3/internal/telemetry/domain/repository"
)

type TelemetryUseCase interface {
	CreatePhotoUseCase(input dto.CreatePhotoInputDTO) (dto.CreatePhotoOutputDTO, error)
	CreateGPSUseCase(input dto.CreateGPSInputDTO) (dto.CreateGPSOutputDTO, error)
	CreateGyroscopeUseCase(input dto.CreateGyroscopeInputDTO) (dto.CreateGyroscopeOutputDTO, error)
}

func NewTelemetryUseCaseRegistry(
	gpsRepo repository.GPSRepository,
	gyroscopeRepo repository.GyroscopeRepository,
	photoRepo repository.PhotoRepository,
	storagePort port.StoragePort,
	recognizerPort port.RecognizerPort,
) TelemetryUseCaseRegistry {
	return TelemetryUseCaseRegistry{
		gpsRepo:        gpsRepo,
		gyroscopeRepo:  gyroscopeRepo,
		photoRepo:      photoRepo,
		storagePort:    storagePort,
		recognizerPort: recognizerPort,
	}
}

type TelemetryUseCaseRegistry struct {
	gpsRepo        repository.GPSRepository
	gyroscopeRepo  repository.GyroscopeRepository
	photoRepo      repository.PhotoRepository
	storagePort    port.StoragePort
	recognizerPort port.RecognizerPort
}

package usecase

import (
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/common/log"
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func (r *TelemetryUseCaseRegistry) CreateGPSUseCase(input dto.CreateGPSInputDTO) (dto.CreateGPSOutputDTO, error) {
	gps, err := entity.NewGPS(entity.GPSProps{
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	})
	if err != nil {
		return dto.CreateGPSOutputDTO{}, err
	}

	if err := r.gpsRepo.Store(*gps); err != nil {
		log.InternalErrLog(
			"CreateGPSUseCase",
			"Store GPS to repository",
			err,
		)

		return dto.CreateGPSOutputDTO{}, custom_err.NewInternalErr()
	}

	return dto.CreateGPSOutputDTO{
		GPS: *gps,
	}, nil
}

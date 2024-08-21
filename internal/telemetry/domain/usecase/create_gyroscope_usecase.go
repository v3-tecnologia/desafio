package usecase

import (
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/common/log"
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func (r *TelemetryUseCaseRegistry) CreateGyroscopeUseCase(input dto.CreateGyroscopeInputDTO) (dto.CreateGyroscopeOutputDTO, error) {
	gyroscope, err := entity.NewGyroscope(entity.GyroscopeProps{
		XPosition: input.XPosition,
		YPosition: input.YPosition,
		ZPosition: input.ZPosition,
	})
	if err != nil {
		return dto.CreateGyroscopeOutputDTO{}, err
	}

	if err := r.gyroscopeRepo.Store(*gyroscope); err != nil {
		log.InternalErrLog(
			"CreateGyroscopeUseCase",
			"Store Gyroscope to repository",
			err,
		)

		return dto.CreateGyroscopeOutputDTO{}, custom_err.NewInternalErr()
	}

	return dto.CreateGyroscopeOutputDTO{
		Gyroscope: *gyroscope,
	}, nil
}

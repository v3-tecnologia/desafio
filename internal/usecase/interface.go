package usecase

import "github.com/HaroldoFV/desafio/internal/dto"

type CreateGyroscopeUseCaseInterface interface {
	Execute(input dto.CreateGyroscopeInputDTO) (dto.GyroscopeOutputDTO, error)
}

type CreateGPSUseCaseInterface interface {
	Execute(input dto.CreateGPSInputDTO) (dto.GPSOutputDTO, error)
}

type CreatePhotoUseCaseInterface interface {
	Execute(input dto.CreatePhotoInputDTO) (dto.PhotoOutputDTO, error)
}

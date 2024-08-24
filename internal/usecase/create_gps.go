package usecase

import (
	"github.com/HaroldoFV/desafio/internal/domain"
	"github.com/HaroldoFV/desafio/internal/domain/entity"
	"github.com/HaroldoFV/desafio/internal/dto"
)

type CreateGPSUseCase struct {
	GPSRepository domain.GPSRepositoryInterface
}

func NewCreateGPSUseCase(
	gpsRepository domain.GPSRepositoryInterface,
) *CreateGPSUseCase {
	return &CreateGPSUseCase{
		GPSRepository: gpsRepository,
	}
}

func (c *CreateGPSUseCase) Execute(input dto.CreateGPSInputDTO) (dto.GPSOutputDTO, error) {
	gps, err := entity.NewGPS(
		input.Latitude,
		input.Longitude,
		input.MacAddress,
	)
	if err != nil {
		return dto.GPSOutputDTO{}, err
	}
	if err = c.GPSRepository.Create(gps); err != nil {
		return dto.GPSOutputDTO{}, err
	}
	outputDTO := dto.GPSOutputDTO{
		ID:         gps.GetID(),
		Latitude:   gps.GetLatitude(),
		Longitude:  gps.GetLongitude(),
		Timestamp:  gps.GetTimestamp(),
		MacAddress: gps.GetMACAddress(),
	}
	return outputDTO, nil
}

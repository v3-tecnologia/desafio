package usecase

import (
	"github.com/HaroldoFV/desafio/internal/domain"
	"github.com/HaroldoFV/desafio/internal/domain/entity"
	"github.com/HaroldoFV/desafio/internal/dto"
)

type CreateGyroscopeUseCase struct {
	GyroscopeRepository domain.GyroscopeRepositoryInterface
}

func NewCreateGyroscopeUseCase(
	gyroscopeRepository domain.GyroscopeRepositoryInterface,
) *CreateGyroscopeUseCase {
	return &CreateGyroscopeUseCase{
		GyroscopeRepository: gyroscopeRepository,
	}
}

func (c *CreateGyroscopeUseCase) Execute(input dto.CreateGyroscopeInputDTO) (dto.GyroscopeOutputDTO, error) {
	gyroscope, err := entity.NewGyroscope(
		input.Name,
		input.Model,
		input.X,
		input.Y,
		input.Z,
		input.MacAddress,
	)
	if err != nil {
		return dto.GyroscopeOutputDTO{}, err
	}
	if err = c.GyroscopeRepository.Create(gyroscope); err != nil {
		return dto.GyroscopeOutputDTO{}, err
	}
	outputDTO := dto.GyroscopeOutputDTO{
		ID:         gyroscope.GetID(),
		Name:       gyroscope.GetName(),
		Model:      gyroscope.GetModel(),
		X:          gyroscope.GetX(),
		Y:          gyroscope.GetY(),
		Z:          gyroscope.GetZ(),
		Timestamp:  gyroscope.GetTimestamp(),
		MacAddress: gyroscope.GetMACAddress(),
	}
	return outputDTO, nil
}

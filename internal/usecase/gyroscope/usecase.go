package gyroscope

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/gyroscope"
	"github.com/kevenmiano/v3/internal/usecase/gyroscope/create"
)

type UseCase interface {
	Execute(d *domain.Gyroscope) (*domain.Gyroscope, error)
}

func NewCreateGyroscopeUseCase(gyroscopeRepository gyroscope.Repository) UseCase {
	return &create.NewGyroscopeUseCase{
		GyroscopeRepository: gyroscopeRepository,
	}
}

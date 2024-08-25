package gyroscope

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
	"github.com/kevenmiano/v3/internal/usecase/gyroscope"
)

type Service struct {
	logger                 infra.Logger
	createGyroscopeUseCase gyroscope.UseCase
}

type IService interface {
	Create(d *domain.Gyroscope) (*domain.Gyroscope, error)
}

func NewGyroscopeService(logger infra.Logger, createGyroscopeUseCase gyroscope.UseCase) *Service {
	return &Service{
		logger:                 logger,
		createGyroscopeUseCase: createGyroscopeUseCase,
	}
}

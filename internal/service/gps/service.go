package gps

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
	usecase "github.com/kevenmiano/v3/internal/usecase/gps"
)

type Service struct {
	logger           infra.Logger
	createGpsUseCase usecase.UseCase
}

type IService interface {
	Create(d *domain.GPS) (*domain.GPS, error)
}

func NewGpsService(logger infra.Logger, createGpsUseCase usecase.UseCase) *Service {
	return &Service{
		logger:           logger,
		createGpsUseCase: createGpsUseCase,
	}
}

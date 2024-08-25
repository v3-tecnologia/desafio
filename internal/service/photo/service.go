package photo

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
	"github.com/kevenmiano/v3/internal/usecase/photo"
)

type Service struct {
	logger             infra.Logger
	createPhotoUseCase photo.UseCase
	findPhotoUseCase   photo.UseCase
}

type IService interface {
	Create(d *domain.Photo) (*domain.Photo, error)
	Find(d *domain.Photo) (*domain.Photo, error)
}

func NewPhotoService(logger infra.Logger, createPhotoUseCase photo.UseCase, findPhotoUseCase photo.UseCase) *Service {
	return &Service{
		logger:             logger,
		createPhotoUseCase: createPhotoUseCase,
		findPhotoUseCase:   findPhotoUseCase,
	}
}

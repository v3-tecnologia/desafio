package photo

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/photo"
	"github.com/kevenmiano/v3/internal/usecase/photo/create"
	"github.com/kevenmiano/v3/internal/usecase/photo/find"
)

type UseCase interface {
	Execute(d *domain.Photo) (*domain.Photo, error)
}

func NewCreatePhotoUseCase(photoRepository photo.Repository) UseCase {
	return &create.PhotoNewUseCase{
		PhotoRepository: photoRepository,
	}
}

func NewFindPhotoUseCase(photoRepository photo.Repository) UseCase {
	return &find.UseCaseFindPhoto{
		PhotoRepository: photoRepository,
	}
}

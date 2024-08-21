package inmemory_repository

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

func NewPhotoInMemoryRepository() *PhotoInMemoryRepository {
	return &PhotoInMemoryRepository{
		Items: []entity.Photo{},
	}
}

type PhotoInMemoryRepository struct {
	Items []entity.Photo
}

func (r *PhotoInMemoryRepository) Store(photo entity.Photo) error {
	r.Items = append(r.Items, photo)
	return nil
}

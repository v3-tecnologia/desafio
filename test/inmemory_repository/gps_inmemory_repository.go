package inmemory_repository

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

func NewGPSInMemoryRepository() *GPSInMemoryRepository {
	return &GPSInMemoryRepository{
		Items: []entity.GPS{},
	}
}

type GPSInMemoryRepository struct {
	Items []entity.GPS
}

func (r *GPSInMemoryRepository) Store(gps entity.GPS) error {
	r.Items = append(r.Items, gps)
	return nil
}

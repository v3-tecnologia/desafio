package inmemory_repository

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

func NewGyroscopeInMemoryRepository() *GyroscopeInMemoryRepository {
	return &GyroscopeInMemoryRepository{
		Items: []entity.Gyroscope{},
	}
}

type GyroscopeInMemoryRepository struct {
	Items []entity.Gyroscope
}

func (r *GyroscopeInMemoryRepository) Store(gyroscope entity.Gyroscope) error {
	r.Items = append(r.Items, gyroscope)
	return nil
}

package gyroscope

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
)

type Repository interface {
	Create(d *domain.Gyroscope) (*domain.Gyroscope, error)
}

type Database = *infra.Dynamo[domain.Gyroscope]

type IRepository struct {
	database Database
}

func NewGyroscopeRepository(database Database) *IRepository {
	return &IRepository{
		database: database,
	}
}

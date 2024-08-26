package gps

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
)

type Repository interface {
	Create(gps *domain.GPS) (*domain.GPS, error)
}

type Database = *infra.Dynamo[domain.GPS]

type IRepository struct {
	logger   infra.Logger
	database Database
}

func NewGPSRepository(database Database) *IRepository {
	return &IRepository{
		database: database,
	}
}

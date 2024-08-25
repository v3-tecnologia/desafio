package rekognition

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
)

type Repository interface {
	Find(d *domain.Faces) (bool, error)
	Create(d *domain.Faces) (*domain.Faces, error)
}

type IRepository struct {
	database infra.Rekognition
}

func NewRekognitionRepository(database infra.Rekognition) *IRepository {
	return &IRepository{
		database: database,
	}
}

package photo

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/rekognition"
	"github.com/kevenmiano/v3/internal/repository/s3"
)

type Repository interface {
	Find(d *domain.Photo) (bool, error)
	Create(d *domain.Photo) (*domain.Photo, error)
}

type IRepository struct {
	s3Service             s3.Repository
	rekognitionRepository rekognition.Repository
}

func NewPhotoRepository(s3Service s3.Repository, rekognitionRepository rekognition.Repository) *IRepository {
	return &IRepository{
		s3Service:             s3Service,
		rekognitionRepository: rekognitionRepository,
	}
}

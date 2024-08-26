package s3

import (
	"github.com/kevenmiano/v3/internal/domain"
)

type Repository interface {
	Upload(d *domain.Photo) (*domain.Photo, error)
}

type IRepository struct {
}

func NewS3Repository() *IRepository {
	return &IRepository{}
}

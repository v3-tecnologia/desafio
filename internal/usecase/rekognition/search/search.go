package search

import (
	"errors"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/rekognition"
)

type FaceImageUseCase struct {
	RekognitionRepository rekognition.Repository
}

var (
	ErrFaceNotFound = errors.New("face not found")
)

func (f FaceImageUseCase) Execute(d *domain.Faces) (*domain.Faces, error) {

	_, err := f.RekognitionRepository.Find(d)

	if err != nil {
		return nil, ErrFaceNotFound

	}

	return d, nil
}

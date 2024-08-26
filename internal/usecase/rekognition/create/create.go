package create

import (
	"errors"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/rekognition"
)

type IndexFaceUseCase struct {
	RekognitionRepository rekognition.Repository
}

var (
	ErrFaceNotCreated = errors.New("face not created")
)

func (c IndexFaceUseCase) Execute(d *domain.Faces) (*domain.Faces, error) {

	indexed, err := c.RekognitionRepository.Create(d)

	if err != nil {
		return nil, ErrFaceNotCreated
	}

	return indexed, nil
}

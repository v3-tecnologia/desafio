package rekognition

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/rekognition"
	"github.com/kevenmiano/v3/internal/usecase/rekognition/create"
	"github.com/kevenmiano/v3/internal/usecase/rekognition/search"
)

type UseCase interface {
	Execute(d *domain.Faces) (*domain.Faces, error)
}

func NewSearchFaceImageUseCase(rekognitionRepository rekognition.Repository) UseCase {
	return &search.FaceImageUseCase{
		RekognitionRepository: rekognitionRepository,
	}
}

func NewCreateIndexFaceUseCase(rekognitionRepository rekognition.Repository) UseCase {
	return &create.IndexFaceUseCase{
		RekognitionRepository: rekognitionRepository,
	}
}

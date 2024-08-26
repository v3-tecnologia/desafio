package rekognition

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
	"github.com/kevenmiano/v3/internal/usecase/rekognition"
)

type Service struct {
	logger                 infra.Logger
	createIndexFaceUseCase rekognition.UseCase
	findFaceImageUseCase   rekognition.UseCase
}

type IService interface {
	Create(d *domain.Faces) (*domain.Faces, error)
	Find(d *domain.Faces) (*domain.Faces, error)
}

func NewRekognitionService(logger infra.Logger, createIndexFaceUseCase rekognition.UseCase, findFaceImageUseCase rekognition.UseCase) *Service {
	return &Service{
		logger:                 logger,
		createIndexFaceUseCase: createIndexFaceUseCase,
		findFaceImageUseCase:   findFaceImageUseCase,
	}
}

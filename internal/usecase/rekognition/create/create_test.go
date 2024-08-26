package create_test

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/test/dummy"
	r "github.com/kevenmiano/v3/internal/test/mock/internal_/repository/rekognition"
	"github.com/kevenmiano/v3/internal/usecase/rekognition"
	"github.com/kevenmiano/v3/internal/usecase/rekognition/create"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreateIndexFaceUseCase_Execute(t *testing.T) {
	rekognitionRepository := new(r.Repository)

	newCreateUseCase := rekognition.NewCreateIndexFaceUseCase(rekognitionRepository)

	newPhoto := dummy.Photo()

	facesDto := &domain.FacesDto{
		ObjectKey: newPhoto.GetKey(),
	}

	faces := domain.NewFaces(facesDto)

	rekognitionRepository.On("Create", faces).Return(faces, nil).Once()

	photoCreated, _ := newCreateUseCase.Execute(faces)

	assert.Equal(t, faces, photoCreated)

	rekognitionRepository.AssertExpectations(t)

}

func TestNewCreateIndexFaceUseCase_Execute_Error(t *testing.T) {
	rekognitionRepository := new(r.Repository)

	newCreateUseCase := rekognition.NewCreateIndexFaceUseCase(rekognitionRepository)

	newPhoto := dummy.Photo()

	facesDto := &domain.FacesDto{
		ObjectKey: newPhoto.GetKey(),
	}

	faces := domain.NewFaces(facesDto)

	rekognitionRepository.On("Create", faces).Return(nil, create.ErrFaceNotCreated).Once()

	photoCreated, err := newCreateUseCase.Execute(faces)

	assert.NotNil(t, err)

	assert.Nil(t, photoCreated)

	rekognitionRepository.AssertExpectations(t)

}

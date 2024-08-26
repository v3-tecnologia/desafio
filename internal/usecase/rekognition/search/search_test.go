package search_test

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/test/dummy"
	r "github.com/kevenmiano/v3/internal/test/mock/internal_/repository/rekognition"
	"github.com/kevenmiano/v3/internal/usecase/rekognition"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFaceFindUseCase_Execute(t *testing.T) {
	rekognitionRepository := new(r.Repository)

	newFindUseCase := rekognition.NewSearchFaceImageUseCase(rekognitionRepository)

	newPhoto := dummy.Photo()

	facesDto := &domain.FacesDto{
		ObjectKey: newPhoto.GetKey(),
	}

	faces := domain.NewFaces(facesDto)

	rekognitionRepository.On("Find", faces).Return(true, nil).Once()

	photoCreated, _ := newFindUseCase.Execute(faces)

	assert.Equal(t, faces, photoCreated)

	rekognitionRepository.AssertExpectations(t)

}

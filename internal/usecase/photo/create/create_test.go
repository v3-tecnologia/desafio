package create_test

import (
	"testing"

	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/test/dummy"
	r "github.com/kevenmiano/v3/internal/test/mock/internal_/repository/photo"
	"github.com/kevenmiano/v3/internal/usecase/photo"
	"github.com/stretchr/testify/assert"
)

func TestCreatePhotoUseCase_Execute(t *testing.T) {
	photoRepository := new(r.Repository)

	useCase := photo.NewCreatePhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	photoRepository.On("Create", newPhoto).Return(newPhoto, nil).Once()

	photoCreated, _ := useCase.Execute(newPhoto)

	assert.Equal(t, newPhoto, photoCreated)

	photoRepository.AssertExpectations(t)
}

func TestCreatePhotoUseCase_FileNameRequired_Error(t *testing.T) {

	photoRepository := new(r.Repository)

	useCase := photo.NewCreatePhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.FileName = ""

	expectedErr := domain.ErrFileNameRequiredPhoto

	photoCreated, err := useCase.Execute(newPhoto)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, photoCreated)

	photoRepository.AssertExpectations(t)
}

func TestCreatePhotoUseCase_ContentRequired_Error(t *testing.T) {
	photoRepository := new(r.Repository)

	useCase := photo.NewCreatePhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.Content = nil

	expectedErr := domain.ErrContentPhoto

	photoCreated, err := useCase.Execute(newPhoto)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, photoCreated)

	photoRepository.AssertExpectations(t)
}

func TestCreatePhotoUseCase_ContentTypeInvalid_Error(t *testing.T) {

	photoRepository := new(r.Repository)

	useCase := photo.NewCreatePhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.ContentType = ""

	expectedErr := domain.ErrContentTypeInvalidPhoto

	photoCreated, err := useCase.Execute(newPhoto)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, photoCreated)

	photoRepository.AssertExpectations(t)
}

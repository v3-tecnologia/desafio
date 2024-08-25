package find_test

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/test/dummy"
	r "github.com/kevenmiano/v3/internal/test/mock/internal_/repository/photo"
	"github.com/kevenmiano/v3/internal/usecase/photo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPhotoUseCase_Execute(t *testing.T) {
	photoRepository := new(r.Repository)

	findPhotoUseCase := photo.NewFindPhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	photoRepository.On("Find", newPhoto).Return(true, nil).Once()

	photoFound, err := findPhotoUseCase.Execute(newPhoto)

	assert.Nil(t, err)

	assert.Equal(t, newPhoto, photoFound)

	photoRepository.AssertExpectations(t)

}

func TestFindPhotoUseCase_FileNameRequired_Error(t *testing.T) {

	photoRepository := new(r.Repository)

	findPhotoUseCase := photo.NewFindPhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.FileName = ""

	expectedErr := domain.ErrFileNameRequiredPhoto

	photoFound, err := findPhotoUseCase.Execute(newPhoto)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, photoFound)

	photoRepository.AssertExpectations(t)

}

func TestFindPhotoUseCase_ContentRequired_Error(t *testing.T) {

	photoRepository := new(r.Repository)

	findPhotoUseCase := photo.NewFindPhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.Content = nil

	expectedErr := domain.ErrContentPhoto

	photoFound, err := findPhotoUseCase.Execute(newPhoto)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, photoFound)

	photoRepository.AssertExpectations(t)

}

func TestFindPhotoUseCase_ContentTypeInvalid_Error(t *testing.T) {

	photoRepository := new(r.Repository)

	findPhotoUseCase := photo.NewFindPhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.ContentType = ".pdf"

	expectedErr := domain.ErrContentTypeInvalidPhoto

	photoFound, err := findPhotoUseCase.Execute(newPhoto)

	assert.Equal(t, expectedErr, err)

	assert.Nil(t, photoFound)

	photoRepository.AssertExpectations(t)

}

func TestFindRecognizePhotoUseCase_Execute(t *testing.T) {
	photoRepository := new(r.Repository)

	findPhotoUseCase := photo.NewFindPhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	photoRepository.On("Find", newPhoto).Return(true, nil).Once()

	_, err := findPhotoUseCase.Execute(newPhoto)

	assert.Nil(t, err)

	assert.Equal(t, newPhoto.Recognized, true)

	photoRepository.AssertExpectations(t)

}

func TestFindPhotoUseCase_NotFound_Error(t *testing.T) {

	photoRepository := new(r.Repository)

	findPhotoUseCase := photo.NewFindPhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	photoRepository.On("Find", newPhoto).Return(false, nil).Once()

	photoFound, err := findPhotoUseCase.Execute(newPhoto)

	assert.Nil(t, err)

	assert.Equal(t, photoFound.Recognized, false)

	photoRepository.AssertExpectations(t)

}

func TestCreateJpgPhotoUseCase_Execute(t *testing.T) {
	photoRepository := new(r.Repository)

	createJpgPhotoUseCase := photo.NewCreatePhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.ContentType = ".jpg"

	photoRepository.On("Create", newPhoto).Return(newPhoto, nil).Once()

	createdPhoto, err := createJpgPhotoUseCase.Execute(newPhoto)

	assert.Nil(t, err)

	assert.Equal(t, newPhoto, createdPhoto)

	photoRepository.AssertExpectations(t)

}

func TestCreateJpegPhotoUseCase_Execute(t *testing.T) {
	photoRepository := new(r.Repository)

	createJpegPhotoUseCase := photo.NewCreatePhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.ContentType = ".jpeg"

	photoRepository.On("Create", newPhoto).Return(newPhoto, nil).Once()

	createdPhoto, err := createJpegPhotoUseCase.Execute(newPhoto)

	assert.Nil(t, err)

	assert.Equal(t, newPhoto, createdPhoto)

	photoRepository.AssertExpectations(t)

}

func TestCreatePngPhotoUseCase_Execute(t *testing.T) {
	photoRepository := new(r.Repository)

	createPngPhotoUseCase := photo.NewCreatePhotoUseCase(photoRepository)

	newPhoto := dummy.Photo()

	newPhoto.ContentType = ".png"

	photoRepository.On("Create", newPhoto).Return(newPhoto, nil).Once()

	createdPhoto, err := createPngPhotoUseCase.Execute(newPhoto)

	assert.Nil(t, err)

	assert.Equal(t, newPhoto, createdPhoto)

	photoRepository.AssertExpectations(t)

}

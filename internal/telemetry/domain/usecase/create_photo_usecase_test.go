package usecase

import (
	"fmt"
	"strings"

	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
)

func (s *Suite) Test_CreatePhotoUseCase() {
	s.Run("it should be able to create a photo", func() {
		content := "this is a test file"
		fileName := "test.txt"
		file := strings.NewReader(content)

		input := dto.CreatePhotoInputDTO{
			File:     file,
			FileName: fileName,
		}

		output, err := s.useCase.CreatePhotoUseCase(input)
		s.NoError(err)

		imgURLParts := strings.Split(output.Photo.ImageURL, "_")
		registeredDate := imgURLParts[0]
		savedFile := imgURLParts[1]
		s.Equal(fileName, savedFile)

		fileKey := fmt.Sprintf("%s_%s", registeredDate, fileName)
		_, ok := s.storageAdapter.Files[fileKey]
		s.True(ok)
	})

	s.Run("it should be not able to create a photo if have invalid params", func() {
		content := "this is a test file"
		fileName := ""
		file := strings.NewReader(content)

		input := dto.CreatePhotoInputDTO{
			File:     file,
			FileName: fileName,
		}

		_, err := s.useCase.CreatePhotoUseCase(input)
		s.Error(err)
		s.Equal(custom_err.NewValidationErr("image_url cannot be blank").Error(), err.Error())
	})
}

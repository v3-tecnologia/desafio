package usecase

import (
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/common/log"
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func (r *TelemetryUseCaseRegistry) CreatePhotoUseCase(input dto.CreatePhotoInputDTO) (dto.CreatePhotoOutputDTO, error) {
	photo, err := entity.NewPhoto(entity.PhotoProps{
		Filename: input.FileName,
	})
	if err != nil {
		return dto.CreatePhotoOutputDTO{}, err
	}

	if err := r.storagePort.SaveFile(input.File, photo.ImageURL); err != nil {
		log.InternalErrLog(
			"CreatePhotoUseCase",
			"Photo saving error",
			err,
		)

		return dto.CreatePhotoOutputDTO{}, err
	}

	detections, err := r.recognizerPort.Recognize(photo.ImageURL)
	if err != nil {
		log.InternalErrLog(
			"CreatePhotoUseCase",
			"Recognizer port",
			err,
		)

		return dto.CreatePhotoOutputDTO{}, custom_err.NewInternalErr()
	}

	photo.AmountOfFacesDetected = len(detections)

	if photo.AmountOfFacesDetected > 0 {
		photo.IsRecognized = true

		var sum float64
		for _, item := range detections {
			sum += item.Confidence
		}

		photo.ConfidenceMean = sum / float64(len(detections))
	}

	if err := r.photoRepo.Store(*photo); err != nil {
		log.InternalErrLog(
			"CreatePhotoUseCase",
			"Store Photo to repository",
			err,
		)

		return dto.CreatePhotoOutputDTO{}, custom_err.NewInternalErr()
	}

	return dto.CreatePhotoOutputDTO{
		Photo: *photo,
	}, nil
}

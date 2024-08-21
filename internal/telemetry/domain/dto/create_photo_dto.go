package dto

import "github.com/charmingruby/g3/internal/telemetry/domain/entity"

type CreatePhotoInputDTO struct {
	ImageURL     string
	IsRecognized bool
}

type CreatePhotoOutputDTO struct {
	Photo entity.Photo
}

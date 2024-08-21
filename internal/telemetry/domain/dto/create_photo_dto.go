package dto

import (
	"io"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

type CreatePhotoInputDTO struct {
	File     io.Reader
	FileName string
}

type CreatePhotoOutputDTO struct {
	Photo entity.Photo
}

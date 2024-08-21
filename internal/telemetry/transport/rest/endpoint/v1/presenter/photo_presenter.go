package presenter

import (
	"time"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func DomainPhotoToHTTP(photo entity.Photo) HTTPPhoto {
	return HTTPPhoto{
		ID:           photo.ID,
		ImageURL:     photo.ImageURL,
		IsRecognized: photo.IsRecognized,
		CreatedAt:    photo.CreatedAt,
	}

}

type HTTPPhoto struct {
	ID           string    `json:"id"`
	ImageURL     string    `json:"image_url"`
	IsRecognized bool      `json:"is_recognized"`
	CreatedAt    time.Time `json:"created_at"`
}

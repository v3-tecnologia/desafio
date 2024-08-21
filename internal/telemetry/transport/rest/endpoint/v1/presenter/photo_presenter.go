package presenter

import (
	"time"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func DomainPhotoToHTTP(photo entity.Photo) HTTPPhoto {
	return HTTPPhoto{
		ID:                    photo.ID,
		ImageURL:              photo.ImageURL,
		IsRecognized:          photo.IsRecognized,
		AmountOfFacesDetected: photo.AmountOfFacesDetected,
		ConfidenceMean:        photo.ConfidenceMean,
		CreatedAt:             photo.CreatedAt,
	}

}

type HTTPPhoto struct {
	ID                    string    `json:"id"`
	ImageURL              string    `json:"image_url"`
	IsRecognized          bool      `json:"is_recognized"`
	AmountOfFacesDetected int       `json:"amount_of_faces_detected"`
	ConfidenceMean        float64   `json:"confidence_mean"`
	CreatedAt             time.Time `json:"created_at"`
}

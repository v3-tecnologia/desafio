package mapper

import (
	"time"

	"github.com/charmingruby/g3/internal/telemetry/domain/entity"
)

func DomainPhotoToPostgres(photo entity.Photo) PostgresPhoto {
	return PostgresPhoto{
		ID:                    photo.ID,
		ImageURL:              photo.ImageURL,
		IsRecognized:          photo.IsRecognized,
		AmountOfFacesDetected: photo.AmountOfFacesDetected,
		ConfidenceMean:        photo.ConfidenceMean,
		CreatedAt:             photo.CreatedAt,
	}
}

func PostgresPhotoToDomain(photo PostgresPhoto) entity.Photo {
	return entity.Photo{
		ID:                    photo.ID,
		ImageURL:              photo.ImageURL,
		IsRecognized:          photo.IsRecognized,
		AmountOfFacesDetected: photo.AmountOfFacesDetected,
		ConfidenceMean:        photo.ConfidenceMean,
		CreatedAt:             photo.CreatedAt,
	}
}

type PostgresPhoto struct {
	ID                    string    `json:"id" db:"id"`
	ImageURL              string    `json:"image_url" db:"image_url"`
	IsRecognized          bool      `json:"is_recognized" db:"is_recognized"`
	AmountOfFacesDetected int       `json:"amount_of_faces_detected" db:"amount_of_faces_detected"`
	ConfidenceMean        float64   `json:"confidence_mean" db:"confidence_mean"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
}

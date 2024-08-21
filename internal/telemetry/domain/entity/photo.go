package entity

import (
	"time"

	"github.com/charmingruby/g3/internal/common/core"
)

func NewPhoto(imageURL string, isRecognized bool) *Photo {
	return &Photo{
		ID:           core.NewID(),
		ImageURL:     imageURL,
		IsRecognized: isRecognized,
		CreatedAt:    time.Now(),
	}
}

type Photo struct {
	ID           string
	ImageURL     string
	IsRecognized bool
	CreatedAt    time.Time
}

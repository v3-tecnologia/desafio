package entity

import (
	"time"

	"github.com/charmingruby/g3/internal/common/core"
	"github.com/charmingruby/g3/internal/common/custom_err"
)

type PhotoProps struct {
	ImageURL     string
	IsRecognized bool
}

func NewPhoto(props PhotoProps) (*Photo, error) {
	p := Photo{
		ID:           core.NewID(),
		ImageURL:     props.ImageURL,
		IsRecognized: props.IsRecognized,
		CreatedAt:    time.Now(),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Photo) validate() error {
	if p.ImageURL == "" {
		return custom_err.NewValidationErr(
			custom_err.NewRequiredErrMessage("image_url"),
		)
	}

	return nil
}

type Photo struct {
	ID           string
	ImageURL     string
	IsRecognized bool
	CreatedAt    time.Time
}

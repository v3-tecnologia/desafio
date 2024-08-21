package entity

import (
	"fmt"
	"time"

	"github.com/charmingruby/g3/internal/common/core"
	"github.com/charmingruby/g3/internal/common/custom_err"
)

type PhotoProps struct {
	Filename     string
	IsRecognized bool
}

func NewPhoto(props PhotoProps) (*Photo, error) {
	p := Photo{
		ID:                    core.NewID(),
		ImageURL:              props.Filename,
		IsRecognized:          props.IsRecognized,
		AmountOfFacesDetected: 0,
		ConfidenceMean:        0,
		CreatedAt:             time.Now(),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	p.ImageURL = fmt.Sprintf("%s_%s", time.Now().String(), props.Filename)

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
	ID                    string
	ImageURL              string
	IsRecognized          bool
	AmountOfFacesDetected int
	ConfidenceMean        float64
	CreatedAt             time.Time
}

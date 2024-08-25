package find

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/photo"
)

type UseCaseFindPhoto struct {
	PhotoRepository photo.Repository
}

func (c UseCaseFindPhoto) Execute(d *domain.Photo) (*domain.Photo, error) {

	if valid, err := d.Validate(); !valid {
		return nil, err
	}

	recognized, err := c.PhotoRepository.Find(d)

	if err != nil {
		return nil, err
	}

	if recognized {
		d.Recognize()
	}

	return d, nil
}

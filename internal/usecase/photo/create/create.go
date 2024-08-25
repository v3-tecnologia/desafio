package create

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/photo"
)

type PhotoNewUseCase struct {
	PhotoRepository photo.Repository
}

func (c PhotoNewUseCase) Execute(d *domain.Photo) (*domain.Photo, error) {

	if valid, err := d.Validate(); !valid {
		return nil, err
	}

	_, err := c.PhotoRepository.Create(d)

	if err != nil {
		return nil, err
	}

	return d, nil
}

package create

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/gps"
)

type GpsUseCase struct {
	GpsRepository gps.Repository
}

func (uc *GpsUseCase) Execute(d *domain.GPS) (*domain.GPS, error) {

	if ok, e := d.Validate(); !ok {
		return nil, e
	}

	d, err := uc.GpsRepository.Create(d)

	if err != nil {
		return nil, err
	}

	return d, nil
}

package create

import (
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/repository/gyroscope"
)

type NewGyroscopeUseCase struct {
	GyroscopeRepository gyroscope.Repository
}

func (uc *NewGyroscopeUseCase) Execute(d *domain.Gyroscope) (*domain.Gyroscope, error) {

	if ok, e := d.Validate(); !ok {
		return nil, e
	}

	_, err := uc.GyroscopeRepository.Create(d)

	if err != nil {
		return nil, err
	}

	return d, nil

}

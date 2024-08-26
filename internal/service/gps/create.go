package gps

import (
	"github.com/kevenmiano/v3/internal/domain"
)

func (s *Service) Create(d *domain.GPS) (*domain.GPS, error) {

	create, err := s.createGpsUseCase.Execute(d)

	if err != nil {

		s.logger.Error(err.Error())
		return nil, err

	}

	return create, nil

}

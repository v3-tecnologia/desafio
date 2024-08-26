package photo

import (
	"fmt"
	"github.com/kevenmiano/v3/internal/domain"
)

func (s *Service) Create(d *domain.Photo) (*domain.Photo, error) {

	s.logger.Info(fmt.Sprintf("Creating photo %v", d.FileName))

	photoCreated, err := s.createPhotoUseCase.Execute(d)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error creating photo: %v", err))
		return nil, err
	}

	return photoCreated, nil
}

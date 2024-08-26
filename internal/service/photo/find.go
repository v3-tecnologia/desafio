package photo

import (
	"fmt"

	"github.com/kevenmiano/v3/internal/domain"
)

func (s *Service) Find(d *domain.Photo) (*domain.Photo, error) {

	photoFound, err := s.findPhotoUseCase.Execute(d)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error finding photo: %v", err))
		return nil, err
	}

	return photoFound, nil
}

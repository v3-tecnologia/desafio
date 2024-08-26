package rekognition

import "github.com/kevenmiano/v3/internal/domain"

func (s *Service) Create(d *domain.Faces) (*domain.Faces, error) {

	_, err := s.createIndexFaceUseCase.Execute(d)

	if err != nil {
		return nil, err
	}

	return d, nil
}

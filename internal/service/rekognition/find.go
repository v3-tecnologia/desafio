package rekognition

import "github.com/kevenmiano/v3/internal/domain"

func (s *Service) Find(d *domain.Faces) (*domain.Faces, error) {

	_, err := s.findFaceImageUseCase.Execute(d)

	if err != nil {
		return nil, err
	}

	return d, nil
}

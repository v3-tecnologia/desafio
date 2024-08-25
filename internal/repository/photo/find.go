package photo

import "github.com/kevenmiano/v3/internal/domain"

func (s *IRepository) Find(d *domain.Photo) (bool, error) {

	facesDto := &domain.FacesDto{
		ObjectKey: d.GetKey(),
	}

	faces := domain.NewFaces(facesDto)

	recognized, err := s.rekognitionRepository.Find(faces)

	if err != nil {
		return false, err
	}

	return recognized, nil

}

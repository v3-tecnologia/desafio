package photo

import "github.com/kevenmiano/v3/internal/domain"

func (s *IRepository) Create(d *domain.Photo) (*domain.Photo, error) {

	photoCreated, err := s.s3Service.Upload(d)

	if err != nil {
		return nil, err
	}

	return photoCreated, nil
}

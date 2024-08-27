package service

import (
	"desafio/models"
)

type Service struct {
	Repository IRepository
}

var _ IService = (*Service)(nil)

func NewService(repo IRepository) IService {
	return &Service{Repository: repo}
}

func (s *Service) ProcessGyroscopeData(data models.GyroscopeRequest) error {
	err := s.Repository.InsertGyroscopeData(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessGpsData(data models.GpsRequest) error {
	err := s.Repository.InsertGpsData(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessPhotoData(data models.PhotoRequest) error {
	err := s.Repository.InsertPhotoData(data)
	if err != nil {
		return err
	}

	return nil
}

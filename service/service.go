package service

import (
	"github/desafio/models"
)

type Service struct {
	Repository Repository
}

var _ ProcessData = (*Service)(nil)

func NewService(repo Repository) *Service {
	return &Service{Repository: repo}
}

func (s *Service) ProcessGyroscopeData(data models.Gyroscope) error {
	err := s.Repository.InsertGyroscopeData(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessGPSData(data models.GPS) error {
	err := s.Repository.InsertGPSData(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessPhoto(data models.Photo) error {
	err := s.Repository.InsertPhoto(data)
	if err != nil {
		return err
	}

	return nil
}

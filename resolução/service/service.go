package service

import (
	"desafio/models"
)

// Estrutura que contém a instancia do repositorio e contém os metodos do serviço
type Service struct {
	Repository IRepository
}

// Validação que a estrutura assina a interface
var _ IService = (*Service)(nil)

// Contrutor da estrutura de Service
func NewService(repo IRepository) *Service {
	return &Service{Repository: repo}
}

// Função que trata os dados de Gyroscope
func (s *Service) ProcessGyroscopeData(data models.GyroscopeRequest) error {
	err := s.Repository.InsertGyroscopeData(data)
	if err != nil {
		return err
	}

	return nil
}

// Função que trata os dados de Gps
func (s *Service) ProcessGpsData(data models.GpsRequest) error {
	err := s.Repository.InsertGpsData(data)
	if err != nil {
		return err
	}

	return nil
}

// Função que trata os dados de Photo
func (s *Service) ProcessPhotoData(data models.PhotoRequest) error {
	err := s.Repository.InsertPhotoData(data)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"github/desafio/models"
)

// Estrutura que contém as informações de instância do repositório
type Service struct {
	Repository Repository
}

// Validação qua a estrutura Service assina a interface ProcessData
var _ ProcessData = (*Service)(nil)

// Função que serve como construtor da estrutura de service
func NewService(repo Repository) *Service {
	return &Service{Repository: repo}
}

// Função para processar os dados do giroscópio
func (s *Service) ProcessGyroscopeData(data models.Gyroscope) error {
	err := s.Repository.InsertGyroscopeData(data)
	if err != nil {
		return err
	}

	return nil
}

// Função para processar os dados do GPS
func (s *Service) ProcessGPSData(data models.GPS) error {
	err := s.Repository.InsertGPSData(data)
	if err != nil {
		return err
	}

	return nil
}

// Função para processar os dados da foto
func (s *Service) ProcessPhoto(data models.Photo) error {
	err := s.Repository.InsertPhoto(data)
	if err != nil {
		return err
	}

	return nil
}

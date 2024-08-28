package service

import "github/desafio/models"

// Interface que define os métodos do service
type ProcessData interface {
	ProcessGyroscopeData(data models.Gyroscope) error
	ProcessGPSData(data models.GPS) error
	ProcessPhoto(data models.Photo) error
}

//Interface que define os métodos do repositório (repository)
type Repository interface {
	InsertGyroscopeData(data models.Gyroscope) error
	InsertGPSData(data models.GPS) error
	InsertPhoto(data models.Photo) error
}
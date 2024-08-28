package service

import "github/desafio/models"

type ProcessData interface {
	ProcessGyroscopeData(data models.Gyroscope) error
	ProcessGPSData(data models.GPS) error
	ProcessPhoto(data models.Photo) error
}

type Repository interface {
	InsertGyroscopeData(data models.Gyroscope) error
	InsertGPSData(data models.GPS) error
	InsertPhoto(data models.Photo) error
}
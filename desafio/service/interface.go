package service

import "desafio/models"

type IService interface {
	ProcessGyroscopeData(data models.GyroscopeRequest) error
	ProcessGpsData(data models.GpsRequest) error
	ProcessPhotoData(data models.PhotoRequest) error
}

type IRepository interface {
	InsertGyroscopeData(data models.GyroscopeRequest) error
	InsertGpsData(data models.GpsRequest) error
	InsertPhotoData(data models.PhotoRequest) error
}

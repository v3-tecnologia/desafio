package service

import "desafio/models"

// Interface que define os metodos dos serviço, essa interface será usada também para abstração
type IService interface {
	ProcessGyroscopeData(data models.GyroscopeRequest) error
	ProcessGpsData(data models.GpsRequest) error
	ProcessPhotoData(data models.PhotoRequest) error
}

// Interface que define os metodos dos repositorios, essa interface será usada também para abstração
type IRepository interface {
	InsertGyroscopeData(data models.GyroscopeRequest) error
	InsertGpsData(data models.GpsRequest) error
	InsertPhotoData(data models.PhotoRequest) error
}

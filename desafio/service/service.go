package service

import (
	"desafio/models"
	"desafio/repository"
)

func ProcessGyroscopeData(data models.GyroscopeRequest) error {
	err := repository.InsertGyroscopeData(data)
	if err != nil {
		return err
	}

	return nil
}

func ProcessGpsData(data models.GpsRequest) error {
	err := repository.InsertGpsData(data)
	if err != nil {
		return err
	}

	return nil
}

func ProcessPhotoData(data models.PhotoRequest) error {
	err := repository.InsertPhotoData(data)
	if err != nil {
		return err
	}

	return nil
}

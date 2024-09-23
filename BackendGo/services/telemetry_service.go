package services

import (
    "BackendGo/models"
    "BackendGo/repositories"
)

func SaveGyroscopeData(data models.GyroscopeData) error {
    return repositories.InsertGyroscopeData(data)
}

func SaveGpsData(data models.GpsData) error {
    return repositories.InsertGpsData(data)
}

func SavePhotoData(data models.PhotoData) error {
    return repositories.InsertPhotoData(data)
}

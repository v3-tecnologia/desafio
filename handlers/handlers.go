package handlers

import (
	"encoding/json"
	"github/desafio/models"
	"net/http"
)

func GyroscopeData(w http.ResponseWriter, r *http.Request) {
	var gyroData models.Gyroscope
	json.NewDecoder(r.Body).Decode(&gyroData)

	err := models.ValidateGyroscopeData(&gyroData)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GPSData(w http.ResponseWriter, r *http.Request) {
	var gpsData models.GPS
	json.NewDecoder(r.Body).Decode(&gpsData)

	err := models.ValidateGPSData(&gpsData)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func PhotoData(w http.ResponseWriter, r *http.Request) {
	var photoData models.Photo
	json.NewDecoder(r.Body).Decode(&photoData)

	err := models.ValidatePhotoData(&photoData)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

package handlers

import (
	"encoding/json"
	"github/desafio/models"
	"net/http"

	"log"
)

func GyroscopeData(w http.ResponseWriter, r *http.Request) {
	var gyroData models.Gyroscope
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&gyroData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := models.ValidateGyroscopeData(&gyroData); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GPSData(w http.ResponseWriter, r *http.Request) {
	var gpsData models.GPS

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&gpsData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := models.ValidateGPSData(&gpsData); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func PhotoData(w http.ResponseWriter, r *http.Request) {
	var photoData models.Photo
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&photoData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := models.ValidatePhotoData(&photoData); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

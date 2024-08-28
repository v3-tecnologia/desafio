package handlers

import (
	"encoding/json"
	"github/desafio/models"
	"github/desafio/service"
	"net/http"

	"log"

	"github.com/go-playground/validator/v10"
)

type Handle struct {
	Service service.ProcessData
}

func NewHandle(service *service.Service) *Handle {
	if service == nil {
		return &Handle{}
	}

	return &Handle{Service: service}
}

func (h *Handle) GyroscopeData(w http.ResponseWriter, r *http.Request) {
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

	validate := validator.New()
	if err := validate.Struct(gyroData); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.ProcessGyroscopeData(gyroData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *Handle) GPSData(w http.ResponseWriter, r *http.Request) {
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

	validate := validator.New()
	if err := validate.Struct(gpsData); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.ProcessGPSData(gpsData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *Handle) PhotoData(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(photo); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.ProcessPhoto(photo)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

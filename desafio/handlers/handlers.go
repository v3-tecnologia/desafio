package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type GyroscopeRequest struct {
	Mac           string  `json:"macAddr" validate:"required,mac"`
	X             float64 `json:"x" validate:"required,number"`
	Y             float64 `json:"y" validate:"required,number"`
	Z             float64 `json:"z" validate:"required,number"`
	UnixtimeStamp int64   `json:"timeStamp" validate:"required,number"`
}

type GpsRequest struct {
	Mac           string `json:"macAddr" validate:"required,mac"`
	Lat           string `json:"latitude" validate:"required,latitude"`
	Lon           string `json:"longitude" validate:"required,longitude"`
	UnixtimeStamp int64  `json:"timeStamp" validate:"required,number"`
}

type PhotoRequest struct {
	Mac           string `json:"macAddr" validate:"required,mac"`
	ImageBase64   string `json:"image" validate:"required"`
	UnixtimeStamp int64  `json:"timeStamp" validate:"required,number"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("up and running...")
}

func GyroscopeHandler(w http.ResponseWriter, r *http.Request) {
	var gyroscopeRequest GyroscopeRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := json.NewDecoder(r.Body).Decode(&gyroscopeRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(gyroscopeRequest); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func GpsHandler(w http.ResponseWriter, r *http.Request) {
	var gpsRequest GpsRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := json.NewDecoder(r.Body).Decode(&gpsRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(gpsRequest); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func PhotoHandler(w http.ResponseWriter, r *http.Request) {
	var photoRequest PhotoRequest

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := json.NewDecoder(r.Body).Decode(&photoRequest)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(photoRequest); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
